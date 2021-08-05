package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Cluster struct {
	Id          int          `json:"id"`
	Name        string       `orm:"size(64)" json:"name,omitempty"`
	Description string       `orm:"size(128);null" json:"description,omitempty"`
	State       string       `orm:"size(32);null" json:"state,omitempty"`
	Components  []*Component `orm:"rel(m2m);rel_table(siot_cluster_component)" json:"components,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Cluster))
}
