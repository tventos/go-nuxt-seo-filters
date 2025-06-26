package config

import (
	"fmt"

	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=postgres user=postgres password=postgres dbname=my-shop port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL!")
	}

	fmt.Println("Connected to PostgreSQL")

	DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Filter{}, &models.FilterValue{}, &models.SeoFilterValue{})
}
