package setting

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func RedisSetup()  {
	var err error
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("[error] server setting, fail to parse 'conf/app.ini': %v", err)
	}

	err = cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("[error] server setting, fail to map server 'conf/app.ini' conf: %v", err)
	}
}

