package web

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"time"
	uflog "webgo/pkg/ucloud/log"
	"webgo/web/base"
)

func NewRouter() *gin.Engine  {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.GET("/log:act", handlerSetLogLevel)
	debug := app.Group("debug")
	{
		// Browser open localhost:8080/debug/pprof
		debug.GET("/pprof", pprofHandler(pprof.Index))
		debug.GET("/cmdline", pprofHandler(pprof.Cmdline))
		debug.GET("/profile", pprofHandler(pprof.Profile))
		debug.POST("/symbol", pprofHandler(pprof.Symbol))
		debug.GET("/symbol", pprofHandler(pprof.Symbol))
		debug.GET("/trace", pprofHandler(pprof.Trace))
		debug.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		debug.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		debug.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		debug.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		debug.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		debug.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
	api := app.Group("api")
	{
		api.POST("/*action", handler)
	}
	return app
}

func handlerSetLogLevel(app *gin.Context) {
	switch app.Param("act") {
	case ":debug":
		uflog.SetLogLevel("DEBUG")
		app.String(http.StatusOK, `log Level [DEBUG] set success`)
	case ":info":
		uflog.SetLogLevel("INFO")
		app.String(http.StatusOK, `log Level [INFO] set success`)
	case ":warn":
		uflog.SetLogLevel("WARN")
		app.String(http.StatusOK, `log Level [WARN] set success`)
	case ":error":
		uflog.SetLogLevel("ERROR")
		app.String(http.StatusOK, `log Level [ERROR] set success`)
	case ":panic":
		uflog.SetLogLevel("PANIC")
		app.String(http.StatusOK, `log Level [PANIC] set success`)
	case ":fatal":
		uflog.SetLogLevel("FATAL")
		app.String(http.StatusOK, `log Level [FATAL] set success`)
	default:
		app.String(http.StatusNotFound, "Not Found")
	}
}


func handler(app *gin.Context) {
	var startMillisecond = time.Now().Nanosecond() / 1e6
	defer func() {
		if SHUTTING_DOWN {
			// close connect
		} else {
		}
		app.Header("Keep-Alive", "timeout=3")
	}()
	defer app.Request.Body.Close()

	params := make(map[string]interface{})
	body, _ := ioutil.ReadAll(app.Request.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, &params)
	fmt.Println(err)
	if err != nil {
		// TODO add log
		base.ResponseError(app, "post request body data must be json", 4901, startMillisecond)
	}
	action := base.ParseInterfaceToString(params["Action"])
	if !base.ExistsAPI(action) {
		// TODO add log
		base.ResponseError(app, fmt.Sprintf("API:%s not found", action), 4901, startMillisecond)
	}
	reqId := base.ParseInterfaceToString(params["ReqId"])
	if reqId == "" {
		reqId = base.UUID(16)
	}

	args := base.NewArgs(action)
	for key, value := range params {
		v := base.ParseInterfaceToString(value)
		args.SetRaw(key, v)
	}
	args.ReqId = reqId
	fmt.Println(args)
	uflog.INFOF("%s, RemoteAddr: %s, Api: %s, Begin", args.ReqId, app.Request.RemoteAddr, args.Name)
	resData := base.Call(args)
	base.ResponseSuccess(app, "", 0, resData, startMillisecond)
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(app *gin.Context) {
		handler.ServeHTTP(app.Writer, app.Request)
	}
}