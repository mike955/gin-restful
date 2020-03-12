package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func ServerSetup()  {
	var err error
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("[error] server setting, fail to parse 'conf/app.ini': %v", err)
	}

	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("[error] server setting, fail to map server 'conf/app.ini' conf: %v", err)
	}
}