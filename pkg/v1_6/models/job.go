package models

import "time"

const (
	//JobPending ...
	JobPending string = "pending"
	//JobRunning ...
	JobRunning string = "running"
	//JobError ...
	JobError string = "error"
	//JobStopped ...
	JobStopped string = "stopped"
	//JobFinished ...
	JobFinished string = "finished"
	//JobCanceled ...
	JobCanceled string = "canceled"
	//JobRetrying indicate the job needs to be retried, it will be scheduled to the end of job queue by statemachine after an interval.
	JobRetrying string = "retrying"
	//JobContinue is the status returned by statehandler to tell statemachine to move to next possible state based on trasition table.
	JobContinue string = "_continue"
	// JobScheduled ...
	JobScheduled string = "scheduled"
)

type Job struct {

	// The job ID.
	Id int64 `json:"id,omitempty"`

	// The status of the job.
	Status string `json:"status,omitempty"`

	// The repository handled by the job.
	Repository string `json:"repository,omitempty"`

	// The ID of the policy that triggered this job.
	PolicyId int64 `json:"policy_id,omitempty"`

	// The operation of the job.
	Operation string `json:"operation,omitempty"`

	// The repository's used tag list.
	Tags []string `json:"tags,omitempty"`

	// The creation time of the job.
	CreationTime time.Time `json:"creation_time,omitempty"`

	// The update time of the job.
	UpdateTime time.Time `json:"update_time,omitempty"`
}
