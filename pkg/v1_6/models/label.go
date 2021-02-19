package models

import "time"

//Label v1.6 标签
type Label struct {
	//The ID of label.
	Id int64 `json:"id"`
	//The name of label.
	Name string `json:"name"`
	//The description of label.
	Description string `json:"description"`
	//The color of label.
	Color string `json:"color"`
	//The scope of label, g for global labels and p for project labels.
	Scope string `json:"scope"`
	//The project ID if the label is a project label.
	ProjectId int64 `json:"project_id"`
	//The creation time of label.
	CreationTime time.Time `json:"creation_time"`
	//The update time of label.
	UpdateTime time.Time `json:"update_time"`
	//The label is deleted or not.
	Deleted bool `json:"deleted"`
}
