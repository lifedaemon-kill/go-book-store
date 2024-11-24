package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s",
		os.Getenv("USER"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("SSLMODE"),
		os.Getenv("HOST"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
