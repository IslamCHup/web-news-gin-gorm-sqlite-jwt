package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"web-news-gin-gorm-sqlite-jwt/models"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/mydatabase.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.News{}, &models.Comment{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
