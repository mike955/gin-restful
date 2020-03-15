package controllers

import (
	"fmt"
	"gin-restful/pkg/models"
	"gin-restful/pkg/utils"
	"gin-restful/pkg/utils/library"
	"github.com/gin-gonic/gin"
)

func Login( c *gin.Context)  {
	app := library.NewController(c)
	var account models.Account
	if err := app.Gin.ShouldBindJSON(&account); err != nil {
		fmt.Println(account)
	}
	account, err := models.Login(account.Username, account.Password)
	if err != nil {
		app.LogError(fmt.Sprintf("login error: %s", err))
		app.ResponseFailed(4901, err)
	}

	token, err := utils.GenerateToken(account.Username, account.Password)
	if err != nil {
		app.LogError(fmt.Sprintf("generate token error: %s", err))
		app.ResponseFailed(4901, err)
	}
	_ = app.Rdb.Exists(account.Username)
	//if err != nil {
	//	app.LogError(fmt.Sprintf("set redis error: %s", err))
	//	app.ResponseFailed(4901, err)
	//}
	account.Password = ""
	app.SetHeader("token", token)
	app.ResponseSuccess(0, account)
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

