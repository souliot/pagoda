package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// Operations about Propertys
type PropertyController struct {
	BaseController
}

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

// @router / [get]
func (c *PropertyController) All() {
	variable := c.GetString("variable")
	label := c.GetString("label")
	typ := c.GetString("type")
	limit, _ := c.GetInt("limit", -1)
	from, _ := c.GetInt("from", 0)
	m := &models.Property{
		Variable: variable,
		Label:    label,
		Type:     typ,
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
