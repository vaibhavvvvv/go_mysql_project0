package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:123pass/simplerest?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", "testuser:123pass@tcp(127.0.0.1:3306)/booktable")
	//d, err := gorm.Open("mysql", "vaibhav:vaibhav@tcp(127.0.0.1:3306)/canteen?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:mypassword@tcp(127.0.0.1:3307)/test")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
