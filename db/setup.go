package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"testREST/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test2 password=a sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})

	DB = db
}