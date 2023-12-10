package helper

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetTestDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CleanTestDB() {
	err := os.Remove("../test.db")
	if err != nil {
		panic(err)
	}
}
