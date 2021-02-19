package models

import "time"

//Policy v1.6版本策略
type Policy struct {
	//The policy ID.
	Id int64 `json:"id,omitempty"`
	//The policy name.
	Name string `json:"name"`
	//The description of the policy.
	Description string `json:"description"`
	//The project list that the policy applys to.
	Projects []Project `json:"projects"`
	//The target list.
	Targets []Target `json:"targets"`
	Trigger Trigger  `json:"trigger"`
	//The replication policy filter array.
	Filters []Filter `json:"filters"`
	//Whether to replicate the existing images now.
	ReplicateExistingImageNow bool `json:"replicate_existing_image_now"`
	//Whether to replicate the deletion operation.
	ReplicateDeletion bool `json:"replicate_deletion"`
	//
	TargetProjectName string `json:"target_project_name"`
	//The create time of the policy.
	CreationTime time.Time `json:"creation_time"`
	//The update time of the policy.
	UpdateTime time.Time `json:"update_time"`
	//The error job count number for the policy.
	ErrorJobCount int64 `json:"error_job_count"`
}

//Trigger v1.6版本
type Trigger struct {
	//The replication policy trigger kind. The valid values are manual, immediate and schedule.
	Kind          string         `json:"kind"`
	ScheduleParam *ScheduleParam `json:"schedule_param"`
}

//ScheduleParam v1.6版本
type ScheduleParam struct {
	//The schedule type.
	Type string `json:"type"`
	//Optional, only used when the type is weedly. The valid values are 1-7.
	Weekday int64 `json:"weekday"`
	//The time offset with the UTC 00:00 in seconds.
	OffTime int64 `json:"offtime"`
}
//Filter v1.6版本镜像过滤器
type Filter struct {
	Kind    string `json:"kind"`
	Value   Value  `json:"value"`
	Pattern string `json:"pattern"`
}
type Tag string
type Name string

type Value struct {
	Name
	Label
	Tag
}

type ReqPolicy struct {
	Policy
	Filters []ReqFilter `json:"filters"`
}

type ReqFilter struct {
	Kind     string      `json:"kind"`
	Value    int64       `json:"value"`
	Pattern  string      `json:"pattern"`
	Metadata interface{} `json:"metadata"`
}

type HarborPolicy struct {
	Id                        int64       `json:"id"`
	Name                      string      `json:"name"`
	Description               string      `json:"description"`
	Trigger                   Trigger     `json:"trigger"`
	ReplicateExistingImageNow bool        `json:"replicate_existing_image_now"`
	ReplicateDeletion         bool        `json:"replicate_deletion"`
	Filters                   []ReqFilter `json:"filters"`
	Projects                  []Project   `json:"projects"`
	Targets                   []Target    `json:"targets"`
	TargetProjectName         string      `json:"target_project_name"`
	CreationTime              time.Time   `json:"creation_time"`
	UpdateTime                time.Time   `json:"update_time"`
	ErrorJobCount             int64       `json:"error_job_count"`
}