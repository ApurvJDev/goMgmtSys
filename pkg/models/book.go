package models

import (
	"github.com/ApurvJDev/goMgmtSys/pkg/config"
	"github.com/jinzhu/gorm" //helps us write query
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

// All functions that are needed to talk to db

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // NewRecord() :- query inside gorm
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book // for returning the list/ slice of books in the db
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book
}

// for update we are going to find delete and then create
