package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	sutil "github.com/souliot/siot-util"
)

type AppUser struct {
	Id         int      `json:"id"`
	Appid      string   `orm:"size(128)" json:"appid"`
	Secret     string   `orm:"size(128)" json:"secret"`
	CreateTime int64    `json:"create_time"`
	Salt       string   `orm:"size(128)" json:"salt"`
	Role       string   `orm:"size(128);null" json:"role,omitempty"`
	User       *User    `orm:"rel(fk)" json:"users"`
	Access     []string `orm:"-" json:"access,omitempty"`
	UserIds    []int    `orm:"-" json:"user_ids,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(AppUser))
}

func (m *AppUser) TableName() string {
	return "appuser"
}
func (m *AppUser) Add(a *User) (err error) {
	salt, err := sutil.GenerateSalt()
	if err != nil {
		return
	}

	pass, err := sutil.GeneratePassHash(a.Password, salt)
	if err != nil {
		return
	}
	now := time.Now()

	m.Appid = sutil.To_md5(a.UserName + now.Format(FormatDateTime))[0:16]
	m.Secret = sutil.To_md5(pass + now.Format(FormatDateTime))
	m.Salt = salt
	m.Role = "admin"
	m.User = a
	m.CreateTime = now.Unix()

	_, _, err = o.ReadOrCreate(m, "User")
	return
}

func (m *AppUser) Delete() (err error) {
	_, err = o.Delete(m)

	return
}

func (m *AppUser) DeleteByUser() (err error) {
	qs := o.QueryTable("AppUser")
	_, err = qs.Filter("User__Id__in", m.UserIds).Delete()
	return
}

func (m *AppUser) GetByUser() (err error, errC *sutil.ControllerError) {
	err = o.Read(m, "User")
	if err != nil {
		errC = sutil.ErrNoRecord
	}
	return
}

func (m *AppUser) Valid() (valid bool) {
	err := o.Read(m, "Appid", "Secret")

	if err != nil {
		return false
	}
	return true
}
