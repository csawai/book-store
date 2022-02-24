package config

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialetcsts/postgres"
)

var (
	db * gorm.db
)
//this is not yet tested. 
func connect () {

	//this db needs to be authenticated properly
	d, err := Open("postgres", "postgres:1234/book")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB () *gorm.DB{
	return db
}