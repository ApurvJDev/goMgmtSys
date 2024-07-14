package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ( // this variable will let other files interact with the DB
	db *gorm.DB
)

func Connect() { // Connect() function helps us to open a connection with the DB

	//	if there is error it will go inside err, otherwise db connection will be stored inside d
	d, err := gorm.Open("mysql", "root:apurvjadhav@34/simplerest?charset=utf8&parseTime=True&loc=local")
	//----------------------------------------------------------- these are some requirements by mySql
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB { // to use in other files
	return db
}
