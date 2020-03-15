package library

import (
	"gin-restful/pkg/emsg"
	"github.com/gin-gonic/gin"
	"time"
)

type Controller struct {
	*Base
	Rdb *RedisConn
}

type Response struct {
	Errno int	`json:"error"`
	Msg string	`json:"msg"`
	Data interface{}	`json:"data"`
	Cost int	`json:"cost"`
}

func NewController(g *gin.Context) *Controller {
	controller := &Controller{
		Base: NewBase(g),
		Rdb: NewRedis(),
	}
	return controller
}

func (app *Controller) ResponseSuccess(errno int, data interface{})  {
	app.Gin.JSON(200, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Base.Cost,
	})
	return
}

func (app *Controller) SetHeader(key string, value string)  {
	app.Gin.Header(key, value)
}

func (app *Controller) ResponseFailed( errno int, data interface{})  {
	app.Gin.JSON(200, Response{
		Errno: errno,
		Msg:   emsg.GetMsg(errno),
		Data:  data,
		Cost:  time.Now().Nanosecond() / 1e6 - app.Base.Cost,
	})
	return
}