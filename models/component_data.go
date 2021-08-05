package models

import "github.com/beego/beego/v2/client/orm"

type ComponentData struct {
	Id       int       `json:"id"`
	Property *Property `orm:"rel(fk)" json:"property,omitempty" description:"组件属性"`
	Value    string    `orm:"size(255);null" json:"value,omitempty" description:"属性值"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(ComponentData))
}
