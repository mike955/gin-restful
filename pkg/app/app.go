package app

import (
	"github.com/gin-gonic/gin"
	"time"
)

type App struct {
	Gin *gin.Context
	Cost int
}

func NewApp(c *gin.Context) *App  {
	app := &App{Gin:c, Cost: time.Now().Nanosecond() / 1e6}
	return app
}