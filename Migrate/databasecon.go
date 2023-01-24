package Migrate

import (
	models "github.com/gormCrud/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IntializeDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=@kash123 dbname=DemoDatabase port=8080 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// Create tables
	db.AutoMigrate(&models.Product{})
	return db
}
