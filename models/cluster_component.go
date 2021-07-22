package models

import "github.com/beego/beego/v2/client/orm"

type Component struct {
	Id            int                 `json:"id"`
	MetaComponent *MetaComponent      `orm:"rel(fk)" json:"meta_Component"`
	Hosts         []*Host             `orm:"rel(m2m);rel_table(siot_component_host)" json:"hosts,omitempty"`
	HostsMap      map[string][]string `orm:"-" json:"hosts"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Component))
}
