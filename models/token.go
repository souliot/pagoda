package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	sutil "github.com/souliot/siot-util"
)

type Token struct {
	Id         int    `json:"id,omitempty"`
	Token      string `json:"token,omitempty"`
	Secret     string `json:"secret,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Token))
}

func NewTokenForApp(appid string, secret string) (t *Token) {
	exp := beego.AppConfig.DefaultInt("TokenExp", 24)
	token, _ := sutil.Create_token(appid, secret, int64(exp))
	t = &Token{
		Token:      token,
		Secret:     secret,
		CreateTime: time.Now().Unix(),
	}
	return
}

func NewTokenForUser(username string, password string) (t *Token) {
	exp := beego.AppConfig.DefaultInt("TokenExp", 24)
	now := time.Now()
	secret := sutil.To_md5(password + now.Format(FormatDateTime))

	token, _ := sutil.Create_token(username, secret, int64(exp))
	t = &Token{
		Token:      token,
		Secret:     secret,
		CreateTime: now.Unix(),
	}
	return
}

func (m *Token) Add() (err error) {
	_, _, err = o.ReadOrCreate(m, "Token")
	return
}
func (m *Token) Delete() (err error) {
	_, err = o.Delete(m, "Token")
	return
}

func (m *Token) Find() (err error) {
	err = o.Read(m, "Token")
	return
}

func (m *Token) Valid() (valid bool) {
	err := m.Find()
	if err != nil {
		return false
	}

	_, err = sutil.Token_auth(m.Token, m.Secret)
	if err != nil {
		return false
	}

	return true
}
