package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// 组件管理
type MetaComponentController struct {
	BaseController
}

// @Summary 获取组件信息
// @Description 获取单个组件详细信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [get]
func (c *MetaComponentController) One() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.MetaComponent{
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

// @Summary 获取组件列表
// @Description 获取所有组件列表
// @Param   name         query    string  false        "name"
// @Param   version			query     string  false        "version"
// @Param   page         query    int     false        "page"
// @Param   pageSize     query    int     false        "pageSize"
// @Success 200 {object} doc.ApiResponse
// @router / [get]
func (c *MetaComponentController) All() {
	m := &models.MetaComponent{}
	m.Name = c.GetString("name")
	m.Version = c.GetString("version")
	m.Page, _ = c.GetInt("Page", 1)
	m.PageSize, _ = c.GetInt("PageSize", -1)

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

// @Summary 添加组件
// @Description 添加组件信息
// @Param   body     body    models.MetaComponent  true   "MetaComponent"
// @Success 200 {object} doc.ApiResponse
// @router / [post]
func (c *MetaComponentController) Add() {
	m := &models.MetaComponent{}
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

// @Summary 修改组件
// @Description 修改组件信息
// @Param   body     body    models.MetaComponent  true   "MetaComponent"
// @Success 200 {object} doc.ApiResponse
// @router / [put]
func (c *MetaComponentController) Update() {
	m := &models.MetaComponent{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, m)
	if err != nil {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	err, errC := m.Update()
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

// @Summary 删除单个组件
// @Description 删除单个组件信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [delete]
func (c *MetaComponentController) Delete() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.MetaComponent{
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

// @Summary 批量删除组件
// @Description 批量删除组件
// @Param   ids     query    int  true        "ids"
// @Param   ids     query    int  true        "ids"
// @Success 200 {object} doc.ApiResponse
// @router / [delete]
func (c *MetaComponentController) DeleteMulti() {
	as := c.GetStrings("ids")
	logs.Info(as)
	ids := c.GetIntSlice("ids")
	logs.Info(ids)
	if len(ids) <= 0 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.MetaComponent{
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
