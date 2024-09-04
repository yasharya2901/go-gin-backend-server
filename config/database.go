package config

import (
	"awesomeProject/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Auto migrate the Product model
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate the database: ", err)
	}
}
