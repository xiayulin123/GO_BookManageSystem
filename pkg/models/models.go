package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xiayulin123/Go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name"`
	Author      string `gorm:"json:"author"`
	Publication string `gorm:"json:"publication"`
}

func init() {
	config.Connection()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
