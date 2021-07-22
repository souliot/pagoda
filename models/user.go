package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type User struct {
	PageQuery  `orm:"-"`
	Id         int    `json:"id"`
	LoginName  string `json:"login_name"`
	LoginPwd   string `json:"login_pwd"`
	UserName   string `orm:"size(64);null" json:"username,omitempty"`
	CreateTime int64  `json:"create_time"`
	Ids        []int  `orm:"-" json:"ids,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(User))
}

func (m *User) One() (err error, errC *sutil.ControllerError) {
	err = o.Read(m)
	if err != nil {
		errC = sutil.ErrDbRead
		return
	}
	return
}

func (m *User) All() (ls *List, err error, errC *sutil.ControllerError) {
	ls = new(List)
	ss := make([]*User, 0)
	qs := o.QueryTable(&User{})
	if m.UserName != "" {
		qs = qs.Filter("UserName__iexact", m.UserName)
	}
	if m.LoginName != "" {
		qs = qs.Filter("LoginName__iexact", m.LoginName)
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

func (m *User) ReadOrAdd() (err error, errC *sutil.ControllerError) {
	m.CreateTime = time.Now().Unix()
	_, _, err = o.ReadOrCreate(m, "LoginName")
	if err != nil {
		errC = sutil.ErrDbInsert
		return
	}
	return
}

func (m *User) Add() (err error, errC *sutil.ControllerError) {
	exist := o.QueryTable(&User{}).Filter("LoginName", m.LoginName).Exist()
	if exist {
		err = fmt.Errorf(sutil.ErrDupRecord.Message)
		errC = sutil.ErrDupRecord
		return
	}
	m.CreateTime = time.Now().Unix()
	_, err = o.Insert(m)
	if err != nil {
		errC = sutil.ErrDbInsert
		return
	}
	return
}

func (m *User) Delete() (err error, errC *sutil.ControllerError) {
	_, err = o.Delete(m)
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}

func (m *User) DeleteMulti() (err error, errC *sutil.ControllerError) {
	_, err = o.QueryTable(&User{}).Filter("Id__in", m.Ids).Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}
