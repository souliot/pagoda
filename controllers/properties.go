package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// 组件属性
type PropertyController struct {
	BaseController
}

// @Summary 获取组件属性信息
// @Description 获取单个组件属性详细信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [get]
func (c *PropertyController) One() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Property{
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

// @Summary 获取组件属性列表
// @Description 获取所有组件属性列表
// @Param   variable     query    string  false        "variable"
// @Param   label				 query    string  false        "label"
// @Param   type				 query    string  false        "type"
// @Param   page         query    int     false        "page"
// @Param   pageSize     query    int     false        "pageSize"
// @Success 200 {object} doc.ApiResponse
// @router / [get]
func (c *PropertyController) All() {
	m := &models.Property{}
	m.Variable = c.GetString("variable")
	m.Label = c.GetString("label")
	m.Type = c.GetString("type")
	mcid, _ := c.GetInt("meta_component_id", -1)
	m.Page, _ = c.GetInt("page", 1)
	m.PageSize, _ = c.GetInt("pageSize", -1)
	m.MetaComponent = &models.MetaComponent{
		Id: mcid,
	}

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

// @Summary 添加组件属性
// @Description 添加组件属性信息
// @Param   body     body    models.Property  true   "Property"
// @Success 200 {object} doc.ApiResponse
// @router / [post]
func (c *PropertyController) Add() {
	m := &models.Property{}
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

// @Summary 添加组件属性
// @Description 添加组件属性信息
// @Param   body     body    models.Property  true   "Property"
// @Success 200 {object} doc.ApiResponse
// @router / [put]
func (c *PropertyController) Update() {
	m := &models.Property{}
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

// @Summary 删除单个组件属性
// @Description 删除单个组件属性信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [delete]
func (c *PropertyController) Delete() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Property{
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

// @Summary 批量删除组件属性
// @Description 批量删除组件属性
// @Param   ids     query    int  true        "ids"
// @Param   ids     query    int  true        "ids"
// @Success 200 {object} doc.ApiResponse
// @router / [delete]
func (c *PropertyController) DeleteMulti() {
	as := c.GetStrings("ids")
	logs.Info(as)
	ids := c.GetIntSlice("ids")
	logs.Info(ids)
	if len(ids) <= 0 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.Property{
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
