package director_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/cloudfoundry/bosh-init/director"
)

var _ = Describe("Director", func() {
	var (
		director Director
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("LatestRuntimeConfig", func() {
		It("returns latest runtime config if there is at least one", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/runtime_configs", "limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
	{"properties": "first"},
	{"properties": "second"}
]`),
				),
			)

			cc, err := director.LatestRuntimeConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(cc).To(Equal(RuntimeConfig{Properties: "first"}))
		})

		It("returns error if there is no runtime config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/runtime_configs", "limit=1"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			_, err := director.LatestRuntimeConfig()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No runtime config"))
		})

		It("returns error if info response in non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/runtime_configs"), server)

			_, err := director.LatestRuntimeConfig()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding runtime configs: Director responded with non-successful status code"))
		})

		It("returns error if info cannot be unmarshalled", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/runtime_configs"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := director.LatestRuntimeConfig()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding runtime configs: Unmarshaling Director response"))
		})
	})

	Describe("UpdateRuntimeConfig", func() {
		It("updates runtime config", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/runtime_configs"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"text/yaml"},
					}),
					ghttp.RespondWith(http.StatusOK, `{}`),
				),
			)

			err := director.UpdateRuntimeConfig([]byte("config"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if info response in non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/runtime_configs"), server)

			err := director.UpdateRuntimeConfig(nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Updating runtime config: Director responded with non-successful status code"))
		})
	})
})
