package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApiRet struct {
	Action  string `json:"Action"`
	RetCode int    `json:"RetCode"`
	Message string `json:"Message,omitempty"`
	ReqId   string `json:"ReqId"`
}

type Response struct {
	errno int `json:errno`
	msg string `json:"msg"`
	data interface{} `json:"data"`
	cost int `json:"cost"`
}


func ResponseSuccess (app *gin.Context, msg string, errno int, data interface{}, cost int) {
	app.SecureJSON(http.StatusOK, gin.H{
		"errno": errno,
		"msg": msg,
		"data": data,
		"cost": time.Now().Nanosecond() / 1e6 - cost,
	})
}

func ResponseError(app *gin.Context, msg string, errno int, cost int) {
	app.SecureJSON(http.StatusOK, gin.H{
		"errno": errno,
		"msg": msg,
		"cost": time.Now().Nanosecond() / 1e6 - cost,
	})
}

func Err(retCode int, args ...interface{}) {
	if len(args) != 0 {
		panic(ApiRet{RetCode: retCode, Message: fmt.Sprintf(errorMap[retCode], args...)})
		return
	}
	panic(ApiRet{RetCode: retCode, Message: errorMap[retCode]})
}

func Errs(retCode int, str string) {
	panic(ApiRet{RetCode: retCode, Message: str})
}