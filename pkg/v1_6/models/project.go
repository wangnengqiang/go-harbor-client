package models

import "time"

//Project v1.6版本项目
type Project struct {
	//Project ID.
	ProjectId int64 `json:"project_id"`
	//The owner ID of the project always means the creator of the project.
	OwnerId int64 `json:"owner_id"`
	//The name of the project.
	Name string `json:"name"`
	//The creation time of the project.
	CreationTime time.Time `json:"creation_time"`
	//The update time of the project.
	UpdateTime time.Time `json:"update_time"`
	//A deletion mark of the project.
	Deleted bool `json:"deleted"`
	//The owner name of the project.
	OwnerName string `json:"owner_name"`
	//Correspond to the UI about whether the project's publicity is updatable (for UI).
	Togglable bool `json:"togglable"`
	//The role ID of the current user who triggered the API (for UI) .
	CurrentUserRoleId int64 `json:"current_user_role_id"`
	//The number of the repositories under this project.
	RepoCount int64 `json:"repo_count"`
	//The total number of charts under this project.
	ChartCount int64 `json:"chart_count"`
	//The metadata of the project.
	Metadata *ProjectMetadata `json:"metadata"`
}

//ProjectMetadata v1.6版本项目Metadata
type ProjectMetadata struct {
	//The public status of the project. The valid values are "true", "false".
	Public string `json:"public"`
	//Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project. The valid values are "true", "false".
	EnableContentTrust string `json:"enable_content_trust"`
	//Whether prevent the vulnerable images from running. The valid values are "true", "false".
	PreventVulnerableImagesFromRunning string `json:"prevent_vulnerable_images_from_running"`
	//If the vulnerability is high than severity defined here, the images cann't be pulled. The valid values are "negligible", "low", "medium", "high", "critical".
	PreventVulnerableImagesFromRunningSeverity string `json:"prevent_vulnerable_images_from_running_severity"`
	//Whether scan images automatically when pushing. The valid values are "true", "false".
	AutomaticallyScanImagesOnPush string `json:"automatically_scan_images_on_push"`
}
