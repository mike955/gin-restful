package app

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Setup()  {
	ResponseSetup()
	LogSetup()
}

type App struct {
	Gin *gin.Context
	Cost int
	LogPath string
}

func NewApp(c *gin.Context) *App  {
	app := &App{Gin:c, Cost: time.Now().Nanosecond() / 1e6}
	return app
}