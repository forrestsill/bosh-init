package director_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/cloudfoundry/bosh-init/director"
	"strconv"
)

var _ = Describe("Director", func() {
	var (
		director Director
		server   *ghttp.Server
	)

	BeforeEach(func() {
		director, server = BuildServer()
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.RespondWith(http.StatusOK, `[
					{
					  "id": "1",
					  "timestamp": 1440318199,
					  "user": "fake-user-1",
					  "action": "fake-action",
					  "object_type": "fake-object-type",
					  "object_name": "fake-object-name",
					  "task": "fake-task",
					  "deployment": "fake-deployment",
					  "instance": "fake-instance",
					  "context": {"fake-context-key":"fake-context-value"}
					},
					{
					  "id": "2",
					  "parent_id": "1",
					  "timestamp": 1440318200,
					  "user": "fake-user-2",
					  "action": "fake-action-2",
					  "object_type": "fake-object-type-2",
					  "object_name": "fake-object-name-2",
					  "task": "fake-task-2",
					  "deployment": "fake-deployment-2",
					  "instance": "fake-instance-2",
					  "context": {}
					}
				]`,
				),
			),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Events", func() {
		It("returns events", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events"),
				),
			)

			_, err := director.Events(EventsFilter{})

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'before-id' option", func() {
			beforeID := "3"
			opts := EventsFilter{BeforeID: &beforeID}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "before_id=3"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'before' option", func() {
			before := strconv.FormatInt(time.Date(2015, time.August, 23, 8, 23, 19, 0, time.UTC).Unix(), 10)
			opts := EventsFilter{Before: &before}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "before_time=1440318199"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'after' option", func() {
			after := strconv.FormatInt(time.Date(2015, time.August, 23, 8, 23, 20, 0, time.UTC).Unix(), 10)
			opts := EventsFilter{After: &after}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "after_time=1440318199"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'deploymentName' option", func() {
			deploymentName := "fake-deployment-2"
			opts := EventsFilter{DeploymentName: &deploymentName}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "deployment=test-bosh-2"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'taskID' option", func() {
			taskID := "fake-task"
			opts := EventsFilter{TaskID: &taskID}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "task=303"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("filters events based on 'instance' option", func() {
			instance := "fake-instance-2"
			opts := EventsFilter{Instance: &instance}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "instance=compilation-6"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns a single event based on multiple options", func() {
			instance := "fake-instance-2"
			deploymentName := "fake-deployment-2"
			opts := EventsFilter{
				Instance:       &instance,
				DeploymentName: &deploymentName,
			}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events", "instance=compilation-6&deployment=test-bosh-2"),
				),
			)

			events, err := director.Events(opts)

			expectedEvents(events)

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns no events based on multiple options", func() {
			instance := "fake-instance-2"
			deploymentName := "fake-deployment"
			opts := EventsFilter{
				DeploymentName: &deploymentName,
				Instance:       &instance,
			}
			server.Reset()
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events"),
					ghttp.RespondWith(http.StatusOK, "[]"),
				),
			)

			events, err := director.Events(opts)

			Expect(events).To(HaveLen(0))

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if response is non-200", func() {
			server.Reset()
			AppendBadRequest(ghttp.VerifyRequest("GET", "/events"), server)

			_, err := director.Events(EventsFilter{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding events: Director responded with non-successful status code"))
		})

		It("returns error if response cannot be unmarshalled", func() {
			server.Reset()
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/events"),
					ghttp.RespondWith(http.StatusOK, ``),
				),
			)

			_, err := director.Events(EventsFilter{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"Finding events: Unmarshaling Director response"))
		})
	})

})

func expectedEvents(events []Event) {
	Expect(events[0].ID()).To(Equal("1"))
	Expect(events[0].Timestamp()).To(Equal(time.Date(2015, time.August, 23, 8, 23, 19, 0, time.UTC)))
	Expect(events[0].User()).To(Equal("fake-user-1"))
	Expect(events[0].Action()).To(Equal("fake-action"))
	Expect(events[0].ObjectType()).To(Equal("fake-object-type"))
	Expect(events[0].ObjectName()).To(Equal("fake-object-name"))
	Expect(events[0].TaskID()).To(Equal("fake-task"))
	Expect(events[0].DeploymentName()).To(Equal("fake-deployment"))
	Expect(events[0].Instance()).To(Equal("fake-instance"))
	Expect(events[0].Context()).To(Equal(map[string]interface{}{"fake-context-key": "fake-context-value"}))

	i := len(events) - 1
	if i >= 0 {
		Expect(events[i].ID()).To(Equal("2"))
		Expect(events[i].ParentID()).To(Equal("1"))
		Expect(events[i].Timestamp()).To(Equal(time.Date(2015, time.August, 23, 8, 23, 20, 0, time.UTC)))
		Expect(events[i].User()).To(Equal("fake-user-2"))
		Expect(events[i].Action()).To(Equal("fake-action-2"))
		Expect(events[i].ObjectType()).To(Equal("fake-object-type-2"))
		Expect(events[i].ObjectName()).To(Equal("fake-object-name-2"))
		Expect(events[i].TaskID()).To(Equal("fake-task-2"))
		Expect(events[i].DeploymentName()).To(Equal("fake-deployment-2"))
		Expect(events[i].Instance()).To(Equal("fake-instance-2"))
		Expect(events[i].Context()).To(Equal(map[string]interface{}{}))
	}
}
