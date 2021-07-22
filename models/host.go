package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type Host struct {
	PageQuery   `orm:"-"`
	Id          int    `json:"id"`
	HostName    string `orm:"size(32)" json:"hostname"`
	IP          string `orm:"size(32);column(ip)" json:"ip"`
	User        string `orm:"size(32);column(user)" json:"user"`
	Passwd      string `orm:"size(32);column(passwd)" json:"passwd"`
	RootPasswd  string `orm:"size(32);column(root_passwd)" json:"root_passwd"`
	Description string `orm:"size(128);null" json:"description,omitempty"`
	Ids         []int  `orm:"-" json:"ids,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Host))
}

func (m *Host) One() (err error, errC *sutil.ControllerError) {
	err = o.Read(m)
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	return
}

func (m *Host) All() (ls *List, err error, errC *sutil.ControllerError) {
	ls = new(List)
	ss := make([]*Host, 0)
	qs := o.QueryTable(&Host{})
	if m.HostName != "" {
		qs = qs.Filter("HostName__iexact", m.HostName)
	}
	if m.IP != "" {
		qs = qs.Filter("IP__iexact", m.IP)
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

func (m *Host) Add() (err error, errC *sutil.ControllerError) {
	exist := o.QueryTable(&Host{}).Filter("IP", m.IP).Exist()
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

func (m *Host) Delete() (err error, errC *sutil.ControllerError) {
	_, err = o.Delete(m)
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}

func (m *Host) DeleteMulti() (err error, errC *sutil.ControllerError) {
	_, err = o.QueryTable(&Host{}).Filter("Id__in", m.Ids).Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}
