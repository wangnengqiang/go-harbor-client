package models

import "time"

//Target v1.6版本目标仓库
type Target struct {
	//The target ID.
	Id int64 `json:"id"`
	//The target address URL string.
	Endpoint string `json:"endpoint"`
	//The target name.
	Name string `json:"name"`
	//The target server username.
	Username string `json:"username"`
	//The target server password.
	Password string `json:"password"`
	//Reserved field.
	Type int `json:"type"`
	//Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure"`
	//The create time of the policy.
	CreationTime time.Time `json:"creation_time"`
	//The update time of the policy.
	UpdateTime time.Time `json:"update_time"`
}

type ReqTarget struct {
	//The target address URL string.
	Endpoint string `json:"endpoint"`
	//The target name.
	Name string `json:"name"`
	//The target server username.
	Username string `json:"username"`
	//The target server password.
	Password string `json:"password"`
	//Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure"`
}