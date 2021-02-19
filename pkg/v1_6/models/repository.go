package models

import "time"

//Repository: 本身是一个仓库，这个仓库里面可以放具体的镜像，是指具体的某个镜像的仓库，比如Tomcat下面有很多个版本的镜像，它们共同组成了Tomcat的Repository。
//Registry: 镜像的仓库，比如官方的是Docker Hub，它是开源的，也可以自己部署一个，Registry上有很多的Repository，Redis、Tomcat、MySQL等等Repository组成了Registry。

//Repository V1.6镜像仓库
type Repository struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	ProjectId    int64     `json:"project_id"`
	Description  string    `json:"description"`
	PullCount    int64     `json:"pull_count"`
	StarCount    int64     `json:"star_count"`
	TagsCount    int64     `json:"tags_count"`
	Labels       []Label   `json:"labels"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}

type SimpleRepository struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	ProjectId int64  `json:"project_id"`
}