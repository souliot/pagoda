package models

import "github.com/beego/beego/v2/client/orm"

type Component struct {
	Id            int                 `json:"id"`
	MetaComponent *MetaComponent      `orm:"rel(fk)" json:"meta_Component,omitempty"`
	Hosts         []*Host             `orm:"rel(m2m);rel_table(siot_component_host)" json:"hosts,omitempty"`
	ComponentData []*ComponentData    `orm:"rel(m2m);rel_table(siot_cluster_component_data)" json:"component_data,omitempty"`
	HostsMap      map[string][]string `orm:"-" json:"hosts,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Component))
}
