package main

import (
	"gin-restful/pkg/app"
	"gin-restful/pkg/models"
	"gin-restful/pkg/setting"
	"gin-restful/routes"
)

func init() {
	setting.Setup()
	models.Setup()
	app.Setup()
}

func main() {

	routersInit := routes.InitRouter()
	//endPoint :=fmt.Sprintf(":%d",setting.ServerSetting.HttpPort)
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	//
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
	routersInit.Run(":8000")
	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler:        routersInit,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//if err := http.ListenAndServe() server.ListenAndServe(); err != nil {
	//	fmt.Println(err)
	//}
	//log.Printf("[info] start http server listening %s", endPoint)
	//fmt.Println("[info] start http server listening %s", endPoint)
	//go func() {
	//	if err := server.ListenAndServe(); err != nil {
	//		fmt.Println(err)
	//	}
	//	log.Printf("[info] start http server listening %s", endPoint)
	//	fmt.Println("[info] start http server listening %s", endPoint)
	//}()

}
