package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connecDB to connect DB
func DatabaseInit() {
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	databaseUrlMysql := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := databaseUrlMysql

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}

	fmt.Println("Database Connected")
	// db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
}
