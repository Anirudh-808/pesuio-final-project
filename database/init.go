package database

import (
	"github.com/anuragrao04/pesuio-final-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(databaseFileName string) (*gorm.DB, error) {
	// implement
	// populate DB variable

	//open the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("database connection screwed up")
	}

	db.AutoMigrate(&models.User{})

	return db, err
}
