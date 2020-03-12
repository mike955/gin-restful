package app

import (
	"gin-restful/pkg/emsg"
	"time"
)

type Response struct {
	Errno int			`json:errno`
	Msg string			`json:msg`
	Data interface{}	`json:data`
	Cost int			`json:cost`
}

func (app *App) ResponseSuccess(httpCode, errno int, data interface{})  {
	app.Gin.JSON(httpCode, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Cost,
	})
	return
}

func (app *App) SetHeader(key string, value string)  {
	app.Gin.Header(key, value)
}

func (app *App) ResponseFailed(httpCode, errno int, data interface{})  {
	app.Gin.JSON(httpCode, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Cost,
	})
	return
}