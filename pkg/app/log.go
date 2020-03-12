package app

import (
	"fmt"
	"gin-restful/pkg/setting"
	"log"
	"os"
	"time"
)

func LogSetup()  {
	logPath := fmt.Sprintf("%s/logs", setting.ServerSetting.LogPath)
	fileName := time.Now().Format("2006-01-02")
	logFile := fmt.Sprintf("%s/%s.log", logPath, fileName)
	res, err := exists(logPath)
	if err != nil {
		log.Fatalf("create dir file error ")
	}
	if res == false {
		os.MkdirAll(logPath,0777)
		file,err:=os.Create(logFile)
		if err!= nil{
			log.Fatalf("create log file error ")
		}
		defer file.Close()
	} else {
		res, err = exists(logFile)
		if err != nil {
			log.Fatalf("create log file error ")
		}
		if res == false {
			file,err:=os.Create(logFile)
			if err!= nil{
				log.Fatalf("create log file error ")
			}
			defer file.Close()
		}
	}
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file) // 将文件设置为log输出的文件
	log.SetPrefix("[INFO]")	// 设置默认前缀
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func (app *App) LogInfo(info string)  {
	log.SetPrefix("[INFO]")
	log.Println(info)
}

func (app *App) LogError(info string)  {
	log.SetPrefix("[ERROR]")
	log.Println(info)
}