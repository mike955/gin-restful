package library

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Base struct {
	Gin *gin.Context
	Cost int
}

func NewBase(g *gin.Context) *Base {
	base := &Base{
		Gin: g,
		Cost: time.Now().Nanosecond() / 1e6,
	}
	return base
}

func (base *Base) LogInfo(info string)  {
	log.SetPrefix("[INFO]")
	log.Println(info)
}

func (base *Base) LogError(info string)  {
	log.SetPrefix("[ERROR]")
	log.Println(info)
}