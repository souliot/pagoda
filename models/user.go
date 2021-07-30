package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type User struct {
	PageQuery  `orm:"-"`
	Id         int      `json:"id" description:"用户ID"`
	UserName   string   `orm:"size(64)" json:"username" description:"登陆名"`
	Password   string   `orm:"size(64)" json:"password" description:"登陆密码"`
	Name       string   `orm:"size(64);null" json:"realName,omitempty" description:"用户姓名"`
	CreateTime int64    `json:"create_time" description:"创建时间"`
	Salt       string   `orm:"size(128)" json:"salt" description:"密码加密"`
	Ids        []int    `orm:"-" json:"ids,omitempty"`
	Roles      []string `orm:"-" json:"roles,omitempty"`
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
	if m.Name != "" {
		qs = qs.Filter("LoginName__iexact", m.Name)
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

func (m *User) Add() (err error, errC *sutil.ControllerError) {
	salt, err := sutil.GenerateSalt()
	if err != nil {
		errC = sutil.Err500
		return
	}
	m.Password, err = sutil.GeneratePassHash(m.Password, salt)
	if err != nil {
		errC = sutil.Err500
		return
	}
	m.CreateTime = time.Now().Unix()
	m.Salt = salt

	created, _, err := o.ReadOrCreate(m, "UserName")
	if err != nil {
		errC = sutil.ErrDbInsert
		return
	}

	if !created {
		err = fmt.Errorf("用户名已存在！")
		errC = sutil.ErrDupRecord
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
	qs := o.QueryTable(&User{})
	if len(m.Ids) > 0 {
		qs = qs.Filter("Id__in", m.Ids)
	}

	if m.UserName != "" {
		qs = qs.Filter("UserName", m.UserName)
	}

	if m.Name != "" {
		qs = qs.Filter("Name", m.Name)
	}

	_, err = qs.Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
		return
	}
	return
}

func (m *User) Login() (token *Token, err error, errC *sutil.ControllerError) {
	pass := m.Password
	err = o.Read(m, "Username")
	if err != nil {
		errC = sutil.ErrNoUser
		return
	}
	hash, err := sutil.GeneratePassHash(pass, m.Salt)
	if err != nil || m.Password != hash {
		err = errors.New("Login Password Error")
		errC = sutil.ErrUserPass
		return
	}

	token = NewTokenForUser(m.UserName, m.Password)

	err = token.Add()
	if err != nil {
		errC = sutil.ErrUserLogin
		return
	}
	return
}

func (m *User) Logout(token string) (err error, errC *sutil.ControllerError) {
	t := &Token{
		Token: token,
	}
	err = t.Delete()
	if err != nil {
		errC = sutil.ErrDbDelete
	}
	return
}

func (m *User) GetUserInfo(token string) (err error, errC *sutil.ControllerError) {
	t := &Token{
		Token: token,
	}
	err = t.Find()
	if err != nil {
		errC = sutil.ErrTokenInValid
		return
	}

	username, err := sutil.Token_auth(t.Token, t.Secret)
	if err != nil {
		errC = sutil.ErrTokenInValid
		return
	}
	m.UserName = username
	err = o.Read(m, "UserName")
	if err != nil {
		errC = sutil.ErrNoUser
		return
	}
	return
}
