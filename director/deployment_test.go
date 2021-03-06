package director_test

import (
	"fmt"
	"net/http"
	"strings"

	semver "github.com/cppforlife/go-semi-semantic/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/cloudfoundry/bosh-init/director"
)

var _ = Describe("Deployment", func() {
	var (
		director   Director
		deployment Deployment
		server     *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()

		var err error

		deployment, err = director.FindDeployment("dep")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Name", func() {
		It("returns name", func() {
			Expect(deployment.Name()).To(Equal("dep"))
		})
	})

	Describe("Releases", func() {
		It("returns releases", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
	{"name": "dep", "releases":[{"name":"rel","version":"ver"}]}
]`),
				),
			)

			rels, err := deployment.Releases()
			Expect(err).ToNot(HaveOccurred())
			Expect(rels[0].Name()).To(Equal("rel"))
			Expect(rels[0].Version()).To(Equal(semver.MustNewVersionFromString("ver")))

			// idempotency check
			rels, err = deployment.Releases()
			Expect(err).ToNot(HaveOccurred())
			Expect(rels[0].Name()).To(Equal("rel"))
			Expect(rels[0].Version()).To(Equal(semver.MustNewVersionFromString("ver")))
		})

		It("returns an error for invalid release versions", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.RespondWith(http.StatusOK, `[
	{"name": "dep", "releases":[{"name":"rel","version":"-"}]}
]`),
				),
			)

			_, err := deployment.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing version for release"))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/deployments"), server)

			_, err := deployment.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Finding deployments"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("GET", "/deployments"), "", server)

			_, err := deployment.Releases()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Finding deployments"))
		})
	})

	Describe("Stemcells", func() {
		It("returns stemcells", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[
	{"name": "dep", "stemcells":[{"name":"rel","version":"ver"}]}
]`),
				),
			)

			stems, err := deployment.Stemcells()
			Expect(err).ToNot(HaveOccurred())
			Expect(stems[0].Name()).To(Equal("rel"))
			Expect(stems[0].Version()).To(Equal(semver.MustNewVersionFromString("ver")))

			// idempotency check
			stems, err = deployment.Stemcells()
			Expect(err).ToNot(HaveOccurred())
			Expect(stems[0].Name()).To(Equal("rel"))
			Expect(stems[0].Version()).To(Equal(semver.MustNewVersionFromString("ver")))
		})

		It("returns an error for invalid stemcell versions", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.RespondWith(http.StatusOK, `[
	{"name": "dep", "stemcells":[{"name":"rel","version":"-"}]}
]`),
				),
			)

			_, err := deployment.Stemcells()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing version for stemcell"))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/deployments"), server)

			_, err := deployment.Stemcells()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Finding deployments"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("GET", "/deployments"), "", server)

			_, err := deployment.Stemcells()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Finding deployments"))
		})
	})

	Describe("Manifest", func() {
		It("returns manifest", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments/dep"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `{"manifest":"content"}`),
				),
			)

			man, err := deployment.Manifest()
			Expect(err).ToNot(HaveOccurred())
			Expect(man).To(Equal("content"))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/deployments/dep"), server)

			_, err := deployment.Manifest()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Fetching manifest"))
		})
	})

	Describe("FetchLogs", func() {
		It("returns logs result", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments/dep/jobs/job/id/logs", "type=job"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				``,
				server,
			)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/tasks/123"),
					ghttp.RespondWith(http.StatusOK, `{"result":"logs-blob-id"}`),
				),
			)

			result, err := deployment.FetchLogs(NewInstanceSlug("job", "id"), nil, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(LogsResult{BlobstoreID: "logs-blob-id", SHA1: ""}))
		})

		It("is able to apply filters and fetch agent logs", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments/dep/jobs/job/id/logs", "type=agent&filters=f1,f2"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				``,
				server,
			)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/tasks/123"),
					ghttp.RespondWith(http.StatusOK, `{"result":"logs-blob-id"}`),
				),
			)

			result, err := deployment.FetchLogs(
				NewInstanceSlug("job", "id"), []string{"f1", "f2"}, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(LogsResult{BlobstoreID: "logs-blob-id", SHA1: ""}))
		})

		It("returns error if task response is non-200", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments/dep/jobs/job/id/logs", "type=job"),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				``,
				server,
			)

			AppendBadRequest(ghttp.VerifyRequest("GET", "/tasks/123"), server)

			_, err := deployment.FetchLogs(NewInstanceSlug("job", "id"), nil, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Finding task '123'"))
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("GET", "/deployments/dep/jobs/job/id/logs", "type=job"), server)

			_, err := deployment.FetchLogs(NewInstanceSlug("job", "id"), nil, false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Fetching logs"))
		})
	})

	Describe("EnableResurrection", func() {
		It("enables resurrection for all instances and returns without an error", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/job/id/resurrection"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"application/json"},
					}),
					ghttp.VerifyBody([]byte(`{"resurrection_paused":false}`)),
				),
			)

			err := deployment.EnableResurrection(NewInstanceSlug("job", "id"), true)
			Expect(err).ToNot(HaveOccurred())
		})

		It("disables resurrection for all instances and returns without an error", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/job/id/resurrection"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"application/json"},
					}),
					ghttp.VerifyBody([]byte(`{"resurrection_paused":true}`)),
				),
			)

			err := deployment.EnableResurrection(NewInstanceSlug("job", "id"), false)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/job/id/resurrection"), server)

			err := deployment.EnableResurrection(NewInstanceSlug("job", "id"), true)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Changing VM resurrection state"))
		})
	})

	Describe("job states", func() {
		var (
			slug  AllOrPoolOrInstanceSlug
			sd    SkipDrain
			force bool
		)

		BeforeEach(func() {
			slug = AllOrPoolOrInstanceSlug{}
			sd = SkipDrain{}
			force = false
		})

		states := map[string]func(Deployment) error{
			"started":  func(d Deployment) error { return d.Start(slug) },
			"detached": func(d Deployment) error { return d.Stop(slug, true, sd, force) },
			"stopped":  func(d Deployment) error { return d.Stop(slug, false, sd, force) },
			"restart":  func(d Deployment) error { return d.Restart(slug, sd, force) },
			"recreate": func(d Deployment) error { return d.Recreate(slug, sd, force) },
		}

		for state, stateFunc := range states {
			state := state
			stateFunc := stateFunc

			Describe(fmt.Sprintf("change state to '%s'", state), func() {
				Describe("with successfully fetched manifest", func() {
					BeforeEach(func() {
						server.AppendHandlers(
							ghttp.CombineHandlers(
								ghttp.VerifyRequest("GET", "/deployments/dep"),
								ghttp.RespondWith(http.StatusOK, `{"manifest":"content"}`),
							),
						)
					})

					It("changes state for specific instance", func() {
						slug = NewAllOrPoolOrInstanceSlug("job", "id")

						ConfigureTaskResult(
							ghttp.CombineHandlers(
								ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/job/id", fmt.Sprintf("state=%s", state)),
								ghttp.VerifyBasicAuth("username", "password"),
								ghttp.VerifyHeader(http.Header{
									"Content-Type": []string{"text/yaml"},
								}),
								ghttp.VerifyBody([]byte("content")),
							),
							``,
							server,
						)

						Expect(stateFunc(deployment)).ToNot(HaveOccurred())
					})

					It("changes state for the whole deployment", func() {
						slug = NewAllOrPoolOrInstanceSlug("", "")

						ConfigureTaskResult(
							ghttp.CombineHandlers(
								ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/*", fmt.Sprintf("state=%s", state)),
								ghttp.VerifyBasicAuth("username", "password"),
								ghttp.VerifyHeader(http.Header{
									"Content-Type": []string{"text/yaml"},
								}),
								ghttp.VerifyBody([]byte("content")),
							),
							``,
							server,
						)

						Expect(stateFunc(deployment)).ToNot(HaveOccurred())
					})

					It("changes state for all indicies of a job", func() {
						slug = NewAllOrPoolOrInstanceSlug("job", "")

						ConfigureTaskResult(
							ghttp.CombineHandlers(
								ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/job", fmt.Sprintf("state=%s", state)),
								ghttp.VerifyBasicAuth("username", "password"),
								ghttp.VerifyHeader(http.Header{
									"Content-Type": []string{"text/yaml"},
								}),
								ghttp.VerifyBody([]byte("content")),
							),
							``,
							server,
						)

						Expect(stateFunc(deployment)).ToNot(HaveOccurred())
					})

					if state != "started" {
						It("changes state with skipping drain and forcing", func() {
							slug = NewAllOrPoolOrInstanceSlug("", "")
							sd = SkipDrain{All: true}
							force = true

							query := fmt.Sprintf("state=%s&skip_drain=*&force=true", state)

							ConfigureTaskResult(
								ghttp.CombineHandlers(
									ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/*", query),
									ghttp.VerifyBasicAuth("username", "password"),
									ghttp.VerifyHeader(http.Header{
										"Content-Type": []string{"text/yaml"},
									}),
									ghttp.VerifyBody([]byte("content")),
								),
								``,
								server,
							)

							Expect(stateFunc(deployment)).ToNot(HaveOccurred())
						})
					}

					It("returns an error if changing state response is non-200", func() {
						AppendBadRequest(ghttp.VerifyRequest("PUT", "/deployments/dep/jobs/*"), server)

						err := stateFunc(deployment)
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring("Changing state"))
					})
				})

				It("returns an error if manifest response is non-200", func() {
					AppendBadRequest(ghttp.VerifyRequest("GET", "/deployments/dep"), server)

					err := stateFunc(deployment)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("Fetching manifest"))
				})
			})
		}
	})

	Describe("ExportRelease", func() {
		var (
			relSlug ReleaseSlug
			osSlug  OSVersionSlug
		)

		BeforeEach(func() {
			relSlug = NewReleaseSlug("rel", "1")
			osSlug = NewOSVersionSlug("os", "2")
		})

		It("returns exported release result", func() {
			reqBody := `{
"deployment_name":"dep",
"release_name":"rel",
"release_version":"1",
"stemcell_os":"os",
"stemcell_version":"2"
}`

			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/releases/export"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"application/json"},
					}),
					ghttp.VerifyBody([]byte(strings.Replace(reqBody, "\n", "", -1))),
				),
				`{"blobstore_id":"release-blob-id","sha1":"release-sha1"}`,
				server,
			)

			result, err := deployment.ExportRelease(relSlug, osSlug)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(ExportReleaseResult{
				BlobstoreID: "release-blob-id",
				SHA1:        "release-sha1",
			}))
		})

		It("returns error if task response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/releases/export"), server)

			_, err := deployment.ExportRelease(relSlug, osSlug)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Exporting release"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("POST", "/releases/export"), ``, server)

			_, err := deployment.ExportRelease(relSlug, osSlug)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unmarshaling export release result"))
		})
	})

	Describe("Update", func() {
		It("succeeds updating deployment", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/deployments", ""),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"text/yaml"},
					}),
					ghttp.VerifyBody([]byte("manifest")),
				),
				``,
				server,
			)

			err := deployment.Update([]byte("manifest"), false, SkipDrain{})
			Expect(err).ToNot(HaveOccurred())
		})

		It("succeeds updating deployment with recreate and skip drain flags", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/deployments", "recreate=true&skip_drain=*"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.VerifyHeader(http.Header{
						"Content-Type": []string{"text/yaml"},
					}),
					ghttp.VerifyBody([]byte("manifest")),
				),
				``,
				server,
			)

			err := deployment.Update([]byte("manifest"), true, SkipDrain{All: true})
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if task response is non-200", func() {
			AppendBadRequest(ghttp.VerifyRequest("POST", "/deployments"), server)

			err := deployment.Update([]byte("manifest"), false, SkipDrain{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Updating deployment"))
		})
	})

	Describe("Delete", func() {
		It("succeeds deleting", func() {
			ConfigureTaskResult(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/deployments/dep", ""),
					ghttp.VerifyBasicAuth("username", "password"),
				),
				``,
				server,
			)

			Expect(deployment.Delete(false)).ToNot(HaveOccurred())
		})

		It("succeeds deleting with force flag", func() {
			ConfigureTaskResult(ghttp.VerifyRequest("DELETE", "/deployments/dep", "force=true"), ``, server)

			Expect(deployment.Delete(true)).ToNot(HaveOccurred())
		})

		It("succeeds even if error occurrs if deployment no longer exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/deployments/dep"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[]`),
				),
			)

			Expect(deployment.Delete(false)).ToNot(HaveOccurred())
		})

		It("returns delete error if listing deployments fails", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/deployments/dep"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			err := deployment.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting deployment 'dep': Director responded with non-successful status code"))
		})

		It("returns delete error if response is non-200 and deployment still exists", func() {
			AppendBadRequest(ghttp.VerifyRequest("DELETE", "/deployments/dep"), server)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/deployments"),
					ghttp.VerifyBasicAuth("username", "password"),
					ghttp.RespondWith(http.StatusOK, `[{"name": "dep"}]`),
				),
			)

			err := deployment.Delete(false)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Deleting deployment 'dep': Director responded with non-successful status code"))
		})
	})
})
