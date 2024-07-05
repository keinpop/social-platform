package db

import (
	"log"
	"mai-platform/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitCompanyDB() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Company{})

	return db
}
