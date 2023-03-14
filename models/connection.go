package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectionDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/crud_gin"))

	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&Product{})
	DB = db

}