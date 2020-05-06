package web

import "net/http"

var SHUTTING_DOWN = false

func Run()  {
	app := NewRouter()

	server := http.Server{
		Addr:              ":8081",
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
}