package db

import (
	"log"
	"os"

	"github.com/gmanninglive/scrapi/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Review{})

	return db
}
