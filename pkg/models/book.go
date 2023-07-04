package models

import (
	"github.com/jinzhu/gorm"
	"github.com/vaibhavvvvv/goMysql/pkg/config"
)

// var db *gorm.DB
// var Book struct {
// 	gorm.Model
// 	Name        string `gorm:""json:"name"` //empty string in gorm:"" defines to use defailt column name i.e. "name"
// 	Author      string `json:"author"`
// 	Publication string `json:"publication"`
// }

// func init() {
// 	config.Connect()
// 	db = config.GetDB()     //getDB in config file returns db
// 	db.AutoMigrate(&Book{}) //creates new table in DB with Book struct and updates if struct changes
// }

// func (b *Book) CreateBook() *Book{
// 	db.NewRecord(b) // checks if model is new in DB or not. returns T or F
// 	db.Create(&b)
// 	return b
// }

// func GetAllBooks()  []Book{
// 	var Books []Book
// 	db.Find(&Books)
// 	return Books
// }

// func GetBookById( Id int64 ) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID=?",Id).Find(&getBook) // ? is a placeholder to be replaced by Id. getBook stores the book with id
// 	return &getBook,db
// }

// func DeleteBookById(Id int64) Book{
// 	var book Book
// 	db.Where("ID=?", Id).Delete(book)
// 	return book
// }

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

//gorm.Model: This field is a pointer to a gorm.Model struct, which is a basic Go struct that
//            includes the following fields: ID, CreatedAt, UpdatedAt, and DeletedAt.
