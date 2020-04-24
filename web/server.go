package web

import (
	"net/http"
	"webgo/web/routers"
)



func Run(zkConn, instanceId, appName string) (err error) {
	// init collect
	//collect.Collection()

	// init router
	app :=  routers.InitRouter()

	// init router

	// use middleware

	// add log

	// load config


	// start server
	server := http.Server{
		Addr:              "",
		Handler:           app,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
	return nil
}