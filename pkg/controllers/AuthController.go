package controllers

import (
	"fmt"
	G "gin-restful/pkg/app"
	"gin-restful/pkg/models"
	"gin-restful/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Login( c *gin.Context)  {
	app := G.NewApp(c)
	//app := NewController(c)
	//app := BaseController{Gin:c, Cost:time.Now().Nanosecond() / 1e6}
	var account models.Account
	if err := app.Gin.ShouldBindJSON(&account); err != nil {
		fmt.Println(account)
	}
	account, err := models.Login(account.Username, account.Password)
	if err != nil {
	}

	token, err := utils.GenerateToken(account.Username, account.Password)
	fmt.Println(token)
	if err != nil {

	}
	account.Password = ""
	app.SetHeader("token", token)
	app.ResponseSuccess(200, 0, account)
}

//func (app *AuthController) Login()  {
//	var account models.Account
//	if err := app.gin.ShouldBindJSON(&account); err != nil {
//		fmt.Println(account)
//	}
//	//var Id, _ = models.Login(account.Username, account.Password)
//	//Id, err := models.Account.Login(&account.Username, &account.Password)
//	fmt.Println("----------------")
//	//fmt.Println(Id)
//	if account.Username != "mike" || account.Password != "18126e7bd3f84b3f3e4df094def5b7de" {
//
//	}
//	//token, err := utils.GenerateToken(account.Username, account.Password)
//	//if err != nil {
//	//
//	//}
//	//c.JSON(200, account)
//	//app.gin.Header("token", token)
//	//app.Response(200, 0, 11, account)
//}

