package entity

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("booking.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Approve{})
	database.AutoMigrate(&Booking{})
	database.AutoMigrate(&Room{})
	database.AutoMigrate(&StatusBook{})
	database.AutoMigrate(&Objective{})

	db = database

}
