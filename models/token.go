package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	sutil "github.com/souliot/siot-util"
)

type Token struct {
	Id         int    `json:"id"`
	Token      string `json:"token"`
	Secret     string `json:"secret"`
	CreateTime int64  `json:"create_time"`
}

func init() {
	orm.RegisterModelWithPrefix(TabPrefix, new(Token))
}

func NewToken(appid string, secret string) (t *Token) {
	exp, _ := beego.AppConfig.Int("TokenExp")
	token, _ := sutil.Create_token(appid, secret, int64(exp))

	t = &Token{
		Token:      token,
		Secret:     secret,
		CreateTime: time.Now().Unix(),
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

	appid, _ := sutil.Token_auth(m.Token, m.Secret)
	valid = o.QueryTable("User").Filter("Appid", appid).Exist()
	return
}
