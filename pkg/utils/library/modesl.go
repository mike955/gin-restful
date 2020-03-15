package library

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DbMap map[string] *gorm.DB

type ModelBase struct {
	ID        uint `gorm:"primary_key" gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time	`gorm:"column:updated_at"`
}

type Model struct {
	//*Base
	Db *gorm.DB
}

func NewModel() *Model  {
	model := &Model{
		//Base: NewBase(g),
		Db: NewMysql(),
	}
	return model
}
