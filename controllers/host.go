package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// 主机管理
type HostController struct {
	BaseController
}

// @Summary 获取主机信息
// @Description 获取单个主机详细信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [get]
func (c *HostController) One() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Host{
		Id: id,
	}

	err, errC := m.One()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.ControllerSuccess{
		Status: 200,
		Data:   m,
	}
	c.ServeJSON()
	return
}

// @Summary 获取主机列表
// @Description 获取所有主机列表
// @Param   hostname     query    string  false        "hostname"
// @Param   ip				   query    string  false        "ip"
// @Param   limit        query    int     false        "limit"
// @Param   from         query    int     false        "from"
// @Success 200 {object} doc.ApiResponse
// @router / [get]
func (c *HostController) All() {
	hostname := c.GetString("hostname")
	ip := c.GetString("ip")
	limit, _ := c.GetInt("limit", -1)
	from, _ := c.GetInt("from", 0)
	m := &models.Host{
		HostName: hostname,
		IP:       ip,
	}

	m.Limit = limit
	m.From = from

	res, err, errC := m.All()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.ControllerSuccess{
		Status: 200,
		Data:   res,
	}
	c.ServeJSON()
	return
}

// @Summary 添加主机
// @Description 添加主机信息
// @Param   body     body    models.Host  true   "Host"
// @Success 200 {object} doc.ApiResponse
// @router / [post]
func (c *HostController) Add() {
	m := &models.Host{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, m)
	if err != nil {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	err, errC := m.Add()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.ControllerSuccess{
		Status: 200,
		Data:   m,
	}
	c.ServeJSON()
	return
}

// @Summary 删除单个主机
// @Description 删除单个主机信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [delete]
func (c *HostController) Delete() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Host{
		Id: id,
	}

	err, errC := m.Delete()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.Actionsuccess
	c.ServeJSON()
	return
}

// @Summary 批量删除主机
// @Description 批量删除主机
// @Param   ids     query    int  true        "ids"
// @Param   ids     query    int  true        "ids"
// @Success 200 {object} doc.ApiResponse
// @router / [delete]
func (c *HostController) DeleteMulti() {
	as := c.GetStrings("ids")
	logs.Info(as)
	ids := c.GetIntSlice("ids")
	logs.Info(ids)
	if len(ids) <= 0 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Host{
		Ids: ids,
	}

	err, errC := m.DeleteMulti()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.Actionsuccess
	c.ServeJSON()
	return
}
