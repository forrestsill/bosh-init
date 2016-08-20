package director

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"net/url"
	"time"
)

type EventImpl struct {
	client Client

	id             string
	parentID       *string
	timestamp      time.Time
	user           string
	action         string
	objectType     string
	objectName     string
	taskID         string
	deploymentName string
	instance       string
	context        map[string]interface{}
}

type EventResp struct {
	ID             string  `json:"id"`
	ParentID       *string `json:"parent_id"`
	Timestamp      int64   `json:"timestamp"`
	User           string  `json:"user"`
	Action         string  `json:"action"`
	ObjectType     string  `json:"object_type"`
	ObjectName     string  `json:"object_name"`
	TaskID         string  `json:"task"`
	DeploymentName string  `json:"deployment"`
	Instance       string  `json:"instance"`
	Context        map[string]interface{}
}

func (e EventImpl) ID() string                      { return e.id }
func (e EventImpl) ParentID() *string               { return e.parentID }
func (e EventImpl) Timestamp() time.Time            { return e.timestamp }
func (e EventImpl) User() string                    { return e.user }
func (e EventImpl) Action() string                  { return e.action }
func (e EventImpl) ObjectType() string              { return e.objectType }
func (e EventImpl) ObjectName() string              { return e.objectName }
func (e EventImpl) TaskID() string                  { return e.taskID }
func (e EventImpl) DeploymentName() string          { return e.deploymentName }
func (e EventImpl) Instance() string                { return e.instance }
func (e EventImpl) Context() map[string]interface{} { return e.context }

func NewEventFromResp(client Client, r EventResp) EventImpl {
	return EventImpl{
		client: client,

		id:             r.ID,
		parentID:       r.ParentID,
		timestamp:      time.Unix(r.Timestamp, 0).UTC(),
		user:           r.User,
		action:         r.Action,
		objectType:     r.ObjectType,
		objectName:     r.ObjectName,
		taskID:         r.TaskID,
		deploymentName: r.DeploymentName,
		instance:       r.Instance,
		context:        r.Context,
	}
}

func (d DirectorImpl) Events(opts EventsFilter) ([]Event, error) {
	events := []Event{}

	eventResps, err := d.client.Events(opts)
	if err != nil {
		return events, err
	}

	for _, r := range eventResps {
		events = append(events, NewEventFromResp(d.client, r))
	}

	return events, nil
}

func (c Client) Events(opts EventsFilter) ([]EventResp, error) {
	var events []EventResp

	u, _ := url.Parse("/events")
	q := u.Query()
	if opts.BeforeID != nil {
		q.Set("before_id", *opts.BeforeID)
	}
	if opts.Before != nil {
		q.Set("before_time", *opts.Before)
	}
	if opts.After != nil {
		q.Set("after_time", *opts.After)
	}
	if opts.DeploymentName != nil {
		q.Set("deployment", *opts.DeploymentName)
	}
	if opts.TaskID != nil {
		q.Set("task", *opts.TaskID)
	}
	if opts.Instance != nil {
		q.Set("instance", *opts.Instance)
	}
	u.RawQuery = q.Encode()

	path := fmt.Sprintf("%v", u)

	err := c.clientRequest.Get(path, &events)
	if err != nil {
		return events, bosherr.WrapErrorf(err, "Finding events")
	}

	return events, nil
}
