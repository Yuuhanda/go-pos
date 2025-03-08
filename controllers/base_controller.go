package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// BaseController defines common methods for all controllers
type BaseController struct {
	beego.Controller
}

// Response is a standard API response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSONResponse returns a standardized JSON response
func (c *BaseController) JSONResponse(status int, message string, data interface{}) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.ServeJSON()
}
