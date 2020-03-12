package main

import (
	"gin-restful/pkg/models"
	"gin-restful/pkg/setting"
	"gin-restful/routes"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {

	routersInit := routes.InitRouter()
	//endPoint :=fmt.Sprintf(":%d",setting.ServerSetting.HttpPort)
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout

	routersInit.Run(":8000")
	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler: routersInit,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//log.Printf("[info] start http server listening %s", endPoint)
	////server.ListenAndServe()
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
