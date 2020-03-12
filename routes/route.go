package routes

import (
	"fmt"
	"gin-restful/pkg/controllers"
	"gin-restful/pkg/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//handle := new(controllers.AuthController)
	r.POST("/getCaptcha", generateCaptcha)
	//r.POST("/session/login", new(controllers.AuthController).Login)

	r.POST("/session/login", controllers.Login )
	//routerHandle("/session/login", "NewAuthController", "login" )
	//api := r.Group("/api")
	//{
	//	//r.POST("/session/login", LoginSession)
	//	r.POST("/session/jwt")
	//}
	return r
}
//
//func routerHandle(path string, controllerName interface{}, funcName string)  {
//	controller := controllers.controllerName.
//	controller[funcName]
//}

type Login struct {
	Username     string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateCaptcha(c *gin.Context)  {
	res := captcha.New()
	fmt.Println(res)
}

func LoginSession(c * gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		fmt.Println(login)
	}
	if login.Username != "mike" || login.Password != "18126e7bd3f84b3f3e4df094def5b7de" {

	}
	token, err := utils.GenerateToken(login.Username, login.Password)
	if err != nil {

	}
	c.Header("token", token)
	c.JSON(200, gin.H{
		"username": login.Username,
	})
}
