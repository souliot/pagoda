package controllers

import (
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// BaseController
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	_, action := c.GetControllerAndAction()
	if action == "Login" || action == "Active" || c.Ctx.Input.IsWebsocket() {
		return
	}
	valid := false
	token := c.Ctx.Input.Header("Authorization")
	if token != "" {
		t := &models.Token{
			Token: token,
		}

		valid = t.Valid()
		if valid {
			return
		}
		t.Delete()
	}

	appid := c.Ctx.Input.Header("Appid")
	secret := c.Ctx.Input.Header("Secret")
	if appid == "" || secret == "" {
		c.Data["json"] = &sutil.ErrExpired
		c.ServeJSON()
		return
	}
	u := &models.AppUser{
		Appid:  appid,
		Secret: secret,
	}
	valid = u.Valid()
	if !valid {
		c.Data["json"] = &sutil.ErrAppid
		c.ServeJSON()
		return
	}
	return
}

func (c *BaseController) GetIntSlice(key string) (res []int) {
	sts := c.GetStrings(key)
	res = make([]int, 0)
	for _, v := range sts {
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}
	return
}

func (c *BaseController) GetInt32Slice(key string) (res []int32) {
	sts := c.GetStrings(key)
	res = make([]int32, 0)
	for _, v := range sts {
		i64, _ := strconv.ParseInt(v, 10, 32)
		res = append(res, int32(i64))
	}
	return
}

func (c *BaseController) GetInt64Slice(key string) (res []int64) {
	sts := c.GetStrings(key)
	res = make([]int64, 0)
	for _, v := range sts {
		i64, _ := strconv.ParseInt(v, 10, 64)
		res = append(res, i64)
	}
	return
}

func (c *BaseController) Options() {
	c.Data["json"] = sutil.Actionsuccess
	c.ServeJSON()
	return
}
