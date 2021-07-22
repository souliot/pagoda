package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// Operations about MetaComponents
type MetaComponentController struct {
	BaseController
}

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

// @router / [get]
func (c *MetaComponentController) All() {
	name := c.GetString("name")
	version := c.GetString("version")
	limit, _ := c.GetInt("limit", -1)
	from, _ := c.GetInt("from", 0)
	m := &models.MetaComponent{
		Name:    name,
		Version: version,
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
