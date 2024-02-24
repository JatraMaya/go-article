package models

import (
	"github.com/jatraMaya/go-library/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Function to initialize Database using SQLite
func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.AppConfig.Database.Name), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize the database")
	}

	DB.AutoMigrate(&Article{}, &User{})
}
