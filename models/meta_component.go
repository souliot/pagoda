package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type MetaComponent struct {
	PageQuery `orm:"-"`
	Id        int       `json:"id" description:"组件ID"`
	Name      string    `orm:"size(64)" json:"name" description:"组件名称"`
	Version   string    `orm:"size(32)" json:"version" description:"组件版本"`
	Property  *Property `orm:"rel(fk)" json:"property" description:"组件属性"`
	Ids       []int     `orm:"-" json:"ids,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(MetaComponent))
}

func (m *MetaComponent) One() (err error, errC *sutil.ControllerError) {
	err = o.Read(m)
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	return
}

func (m *MetaComponent) All() (ls *List, err error, errC *sutil.ControllerError) {
	ls = new(List)
	ss := make([]*MetaComponent, 0)
	qs := o.QueryTable(&MetaComponent{})
	if m.Name != "" {
		qs = qs.Filter("Name__iexact", m.Name)
	}
	if m.Version != "" {
		qs = qs.Filter("Version", m.Version)
	}

	cnt, err := qs.Count()
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	_, err = qs.OrderBy("Id").Limit(m.PageSize, (m.Page-1)*m.PageSize).All(&ss)
	if err != nil {
		errC = sutil.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	ls.Lists = ss
	ls.Total = cnt
	return
}

func (m *MetaComponent) Add() (err error, errC *sutil.ControllerError) {
	exist := o.QueryTable(&MetaComponent{}).Filter("Name", m.Name).Filter("Version", m.Version).Exist()
	if exist {
		err = fmt.Errorf(sutil.ErrDupRecord.Message)
		errC = sutil.ErrDupRecord
		return
	}
	_, err = o.Insert(m)
	if err != nil {
		errC = sutil.ErrDbInsert
		return
	}
	return
}

func (m *MetaComponent) Delete() (err error, errC *sutil.ControllerError) {
	_, err = o.Delete(m)
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}

func (m *MetaComponent) DeleteMulti() (err error, errC *sutil.ControllerError) {
	_, err = o.QueryTable(&MetaComponent{}).Filter("Id__in", m.Ids).Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}
