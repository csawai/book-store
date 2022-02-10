package config

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialetcsts/postgres"
)

var (
	db * gorm.db
)

func connect () {
	d, err := Open("postgres", "chetan:1234/book")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB () *gorm.DB{
	return db
}