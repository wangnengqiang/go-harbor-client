package models

// 镜像中的标签

import "time"

type DetailedTag struct {
	//The digest of the tag.
	Digest string `json:"digest"`
	//The name of the tag.
	Name string `json:"name"`
	//The size of the image.
	Size int64 `json:"size"`
	//The architecture of the image.
	Architecture string `json:"architecture"`
	//The os of the image.
	OS string `json:"os"`
	//The version of docker which builds the image.
	DockerVersion string `json:"docker_version"`
	//The author of the image.
	Author string `json:"author"`
	//The build time of the image.
	Created string `json:"created"`
	//The signature of image, defined by RepoSignature. If it is null, the image is unsigned.
	Signature map[string]interface{} `json:"signature"`
	//The overview of the scan result. This is an optional property.
	ScanOverview InlineModel0 `json:"scan_overview"`
	//The label list.
	Labels []*Label `json:"labels"`
}

type InlineModel0 struct {
	//The digest of the image.
	Digest string `json:"digest"`
	//The status of the scan job, it can be "pendnig", "running", "finished", "error".
	ScanStatus string `json:"scan_status"`
	//The ID of the job on jobservice to scan the image.
	JobId int64 `json:"job_id"`
	//0-Not scanned, 1-Negligible, 2-Unknown, 3-Low, 4-Medium, 5-High.
	Severity int64 `json:"severity"`
	//The top layer name of this image in Clair, this is for calling Clair API to get the vulnerability list of this image.
	DetailsKey string `json:"details_key"`
	//The components overview of the image.
	Components InlineModel1 `json:"components"`
}
type InlineModel1 struct {
	//Total number of the components in this image.
	Total int64 `json:"total"`
	//List of number of components of different severities.
	Summary []ComponentOverviewEntry `json:"summary"`
	//The offest in seconds of UTC 0 o'clock, only valid when the policy type is "daily"
	DailyTime time.Time `json:"daily_time"`
}

type ComponentOverviewEntry struct {
	//1-None/Negligible, 2-Unknown, 3-Low, 4-Medium, 5-High
	Severity int64 `json:"severity"`
	//number of the components with certain severity.
	Count int64 `json:"count"`
}

// TagResp holds the information of one image tag
type TagResp struct {
	DetailedTag
	Signature    *Target          `json:"signature"`
	ScanOverview *ImgScanOverview `json:"scan_overview,omitempty"`
	PushTime     time.Time        `json:"push_time"`
	PullTime     time.Time        `json:"pull_time"`
	Config       *cfg             `json:"config"`
}
type ImgScanOverview struct {
	ID              int64               `orm:"pk;auto;column(id)" json:"-"`
	Digest          string              `orm:"column(image_digest)" json:"image_digest"`
	Status          string              `orm:"-" json:"scan_status"`
	JobID           int64               `orm:"column(scan_job_id)" json:"job_id"`
	Sev             int                 `orm:"column(severity)" json:"severity"`
	CompOverviewStr string              `orm:"column(components_overview)" json:"-"`
	CompOverview    *ComponentsOverview `orm:"-" json:"components,omitempty"`
	DetailsKey      string              `orm:"column(details_key)" json:"details_key"`
	CreationTime    time.Time           `orm:"column(creation_time);auto_now_add" json:"creation_time,omitempty"`
	UpdateTime      time.Time           `orm:"column(update_time);auto_now" json:"update_time,omitempty"`
}
type ComponentsOverview struct {
	Total   int                        `json:"total"`
	Summary []*ComponentsOverviewEntry `json:"summary"`
}
type ComponentsOverviewEntry struct {
	Sev   int `json:"severity"`
	Count int `json:"count"`
}

type cfg struct {
	Labels map[string]string `json:"labels"`
}

//Request to give source image and target tag.
type RetagReq struct {
	//new tag to be created
	Tag string `json:"tag"`
	//Source image to be retagged, e.g. 'stage/app:v1.0'
	SrcImage string `json:"src_image"`
	//If target tag already exists, whether to override it
	Override bool `json:"override"`
}
