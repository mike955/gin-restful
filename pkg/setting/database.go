package setting

import (
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var db *gorm.DB

var DatabaseSetting = &Database{}

func DatabaseSetup()  {
	var err error
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("[error] database setting, fail to parse 'conf/app.ini': %v", err)
	}

	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("[error] database setting, fail to map server 'conf/app.ini' conf: %v", err)
	}
}