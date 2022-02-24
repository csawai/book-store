package models

import (


	"github.com/zinjhu/gorm"
	"github.com/csawai/book-store/pkg/config"

)

var = db "gorm.DB"

type BOOK struct {

	gorm.model
	Name string `gorm: json : "name"`
	Author string `json : " author"`
	Publication string `json : "publication"`
}


func init () {
	config.connect()
	db = config.GetDB()
	db. AutoMigrate(&Book{})
}