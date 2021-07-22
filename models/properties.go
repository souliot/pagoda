package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type Property struct {
	PageQuery   `orm:"-"`
	Id          int    `json:"id"`
	Variable    string `orm:"size(32)" json:"variable"`
	Label       string `orm:"size(32)" json:"label"`
	Description string `orm:"size(128);null" json:"description,omitempty"`
	Type        string `orm:"size(32)" json:"type"`
	Default     string `orm:"size(32);null" json:"default,omitempty"`
	Required    bool   `json:"required"`
	Ids         []int  `orm:"-" json:"ids,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Property))
}

func (m *Property) One() (err error, errC *sutil.ControllerError) {
	err = o.Read(m)
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	return
}

func (m *Property) All() (ls *List, err error, errC *sutil.ControllerError) {
	ls = new(List)
	ss := make([]*Property, 0)
	qs := o.QueryTable(&Property{})
	if m.Type != "" {
		qs = qs.Filter("Type__iexact", m.Type)
	}
	if m.Variable != "" {
		qs = qs.Filter("Variable__iexact", m.Variable)
	}
	if m.Label != "" {
		qs = qs.Filter("Label__iexact", m.Label)
	}
	cnt, err := qs.Count()
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	_, err = qs.OrderBy("Id").Limit(m.Limit, m.From).All(&ss)
	if err != nil {
		errC = sutil.ErrDbRead
		errC.MoreInfo = err.Error()
		return
	}
	ls.Lists = ss
	ls.Total = cnt
	return
}

func (m *Property) Add() (err error, errC *sutil.ControllerError) {
	exist := o.QueryTable(&Property{}).Filter("Type", m.Type).Filter("Variable", m.Variable).Filter("Label", m.Label).Exist()
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

func (m *Property) Delete() (err error, errC *sutil.ControllerError) {
	_, err = o.Delete(m)
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}

func (m *Property) DeleteMulti() (err error, errC *sutil.ControllerError) {
	_, err = o.QueryTable(&Property{}).Filter("Id__in", m.Ids).Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}