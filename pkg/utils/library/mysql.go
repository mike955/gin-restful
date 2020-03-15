package library

import (
	"fmt"
	"gin-restful/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
)

var MysqlConn *gorm.DB

func MysqlSetup() error {
	var err error
	MysqlConn, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	log.Println("[info] database.Setup success: ")
	if setting.ServerSetting.RunMode == "dev" {
		MysqlConn.LogMode(true)
	}
	MysqlConn.SingularTable(true)
	MysqlConn.DB().SetMaxIdleConns(10)
	MysqlConn.DB().SetMaxOpenConns(100)
	return nil
}

func NewMysql() *gorm.DB  {
	return MysqlConn
}