package controllers

import (
	"gin-restful/pkg/emsg"
	"github.com/gin-gonic/gin"
	"time"
)

type BaseController struct {
	// context data
	Gin  *gin.Context
	Cost int
	authType string		// session、jwt
}

type Response struct {
	Errno int			`json:"errno"`
	Msg string			`json:"msg"`
	Data interface{}	`json:"data"`
	Cost int			`json:"cost"`
}

func NewController(c *gin.Context) *BaseController  {
	app := &BaseController{Gin:c, Cost: time.Now().Nanosecond() / 1e6}
	return app
}

func (app *BaseController) ResponseSuccess(httpCode, errno int, data interface{})  {
	app.Gin.JSON(httpCode, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Cost,
	})
	return
}

func (app *BaseController) SetHeader(key string, value string)  {
	app.Gin.Header(key, value)
}

func (app *BaseController) ResponseFailed(httpCode, errno int, data interface{})  {
	app.Gin.JSON(httpCode, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Cost,
	})
	return
}

func (app *BaseController) LogInfo(httpCode, errno int, data interface{})  {
}

func (app *BaseController) LogError(httpCode, errno int, data interface{})  {
}

func (app *BaseController) LogDebug(httpCode, errno int, data interface{})  {
}
