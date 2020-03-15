package models

import (
	"errors"
	"gin-restful/pkg/utils/library"
	"time"
)

type Account struct {
	ID        uint `gorm:"primary_key" gorm:"column:id" json:"id"`
	Username    string `gorm:"column:username" json:"username"`
	Password    string `gorm:"column:password" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createAt"`
	UpdatedAt time.Time	`gorm:"column:updated_at" json:"updateAt"`
}

func Login(username, password string) (Account, error)  {
	app := library.NewModel()
	var account Account
	if err := app.Db.Where("username=?", username).Where("password=?", password).First(&account).Error; err != nil {
		return account, errors.New( "sql error")
	}
	return account, nil
}
