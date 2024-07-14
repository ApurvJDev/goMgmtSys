package models

import (
	"github.com/ApurvJDev/goMgmtSys/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()    // calls config.Connect() to start connection with DB
	db = config.GetDB() // GetDB function returns DB that is stored in db variable
	db.AutoMigrate(&Book{})
}
