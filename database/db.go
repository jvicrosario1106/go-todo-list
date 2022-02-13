package database

import (
	"log"

	"github.com/jvicrosario1106/todo-list/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Print("Database Failed to Connect")

	}

	DB = db

	db.AutoMigrate(&models.Todo{})

	log.Print("Database Successfully Connected!")

}
