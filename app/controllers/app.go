package controllers

import (
	"crud/app/helpers"
	"crud/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var (
	user   models.User
	errors helpers.Error
)

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) parseBody() error {
	err := c.Params.BindJSON(&user)
	return err
}

func (c App) AddUser() revel.Result {
	err := c.parseBody()
	if err != nil {
		c.Response.Status = 400
		errors.Error = "Invalid JSON passed"
		return c.RenderJSON(errors)
	}
	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		errors.Error = errors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(errors)
	}
	response := user.AddUser()
	if response != nil {
		c.Response.Status = 503
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 201
	return nil
}

func (c App) GetUser(id int64) revel.Result {
	data, response := user.GetUser(id)
	if response != nil {
		c.Response.Status = 400
		errors.Error = "User not found"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) UpdateUser(id int64) revel.Result {
	err := c.parseBody()
	if err != nil {
		c.Response.Status = 400
		errors.Error = "Invalid JSON passed"
		return c.RenderJSON(errors)
	}
	user.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		errors.Error = errors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(errors)
	}
	response := user.UpdateUser(id)
	if response != nil {
		c.Response.Status = 503
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) DeleteUser(id int64) revel.Result {
	response := user.DeleteUser
	if response != nil {
		c.Response.Status = 503
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) ListUsers(page int64, perPage int64) revel.Result {
	data, response := user.ListUsers()
	if response != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}
