package models

import "time"

type ClusterLog struct {
	ClusterId string        `json:"cluster_id"`
	Created   time.Time     `json:"created"`
	Log       *Notification `json:"log"`
}

type Notification struct {
	Data  map[string]interface{} `json:"data"`
	Now   string                 `json:"time"`
	Task  *Task                  `json:"task"`
	Stage string                 `json:"stage"`
	State string                 `json:"state"`
	Host  string                 `json:"host"`
}

type Task struct {
	Name  string `json:"name"`
	State string `json:"state"`
}
