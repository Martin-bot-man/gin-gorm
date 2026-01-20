package config

import (
	"fmt"
	"golang-crud-gin/helper"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	// 1. Check for the Render environment variable first
	sqlInfo := os.Getenv("DATABASE_URL")

	// 2. If the variable is empty (like when you run it on your own computer), 
	// fall back to your local settings
	if sqlInfo == "" {
		host := "localhost"
		port := 5432
		user := "postgres"
		password := "postgres"
		dbName := "test"
		sqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
			host, port, user, password, dbName)
	}

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}