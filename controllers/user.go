package controllers

import (
	"encoding/json"

	"github.com/souliot/pagoda/models"
	sutil "github.com/souliot/siot-util"
)

// 用户管理
type UserController struct {
	BaseController
}

// @Summary 获取用户信息
// @Description 获取单个用户详细信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [get]
func (c *UserController) One() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.User{
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

// @Summary 获取用户列表
// @Description 获取所有用户列表
// @Param   username     query    string  false		"username"
// @Param   name         query    string  false   "name"
// @Param   page         query    int     false   "page"
// @Param   pageSize     query    int     false   "pageSize"
// @Success 200 {object} doc.ApiResponse
// @router / [get]
func (c *UserController) All() {
	m := &models.User{}
	m.UserName = c.GetString("username")
	m.Name = c.GetString("name")
	m.Page, _ = c.GetInt("page", 1)
	m.PageSize, _ = c.GetInt("pageSize", -1)

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

// @Summary 添加用户
// @Description 添加用户信息
// @Param   body     body    models.User  true        "User"
// @Success 200 {object} doc.ApiResponse
// @router / [post]
func (c *UserController) Add() {
	m := &models.User{}
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

// @Summary 删除单个用户
// @Description 删除单个用户信息
// @Param   id     path    int  true        "id"
// @Success 200 {object} doc.ApiResponse
// @router /:id [delete]
func (c *UserController) Delete() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.User{
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

// @Summary 批量删除用户
// @Description 批量删除用户
// @Param   ids          query    int     false   "ids"
// @Param   ids     		 query    int     false   "ids"
// @Param   username     query    string  false   "username"
// @Param   name         query    string  false   "name"
// @Success 200 {object} doc.ApiResponse
// @router / [delete]
func (c *UserController) DeleteMulti() {
	ids := c.GetIntSlice("ids")
	username := c.GetString("username")
	name := c.GetString("name")
	if len(ids) <= 0 && username == "" && name == "" {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	m := &models.User{
		Ids:      ids,
		UserName: username,
		Name:     name,
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

// @Summary 用户登录
// @Description 用户登录
// @Param   body     body    models.User  true	"User"
// @Success 200 {object} doc.ApiResponse
// @router /login [post]
func (c *UserController) Login() {
	m := &models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, m)
	if err != nil {
		c.Data["json"] = sutil.ErrUserInput
		c.ServeJSON()
		return
	}
	token, err, errC := m.Login()
	if err != nil {
		errC.MoreInfo = err.Error()
		c.Data["json"] = errC
		c.ServeJSON()
		return
	}
	c.Data["json"] = sutil.ControllerSuccess{
		Status: 200,
		Data:   token,
	}
	c.ServeJSON()
	return
}

// @Summary 用户登出
// @Description 用户登出
// @Success 200 {object} doc.ApiResponse
// @router /logout [post]
func (c *UserController) Logout() {
	token := c.Ctx.Input.Header("Authorization")
	m := &models.User{}

	err, errC := m.Logout(token)
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

// @Summary 获取登录用户信息
// @Description 获取登录用户信息
// @Success 200 {object} doc.ApiResponse
// @router /getUserInfo [get]
func (c *UserController) GetUserInfo() {
	token := c.Ctx.Input.Header("Authorization")
	m := &models.User{}

	err, errC := m.GetUserInfo(token)
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
