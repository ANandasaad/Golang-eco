package database

import (
	"golang-eco/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost port=5432 dbname=test user=postgres password=admin connect_timeout=10"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change the database provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductUser{}, &models.Payment{}, &models.Order{}, &models.Address{})
	if err != nil {
		log.Fatal("Failed to auto migrate models:", err)
	}

}
