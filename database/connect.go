package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Muhammad5943/go-fiber-gorm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connecDB to connect DB
func DatabaseInit() {
	var err error
	// Load .env file
	config.Config(".env")

	// retrieve .env variable
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	// call format database
	databaseUrlMysql := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := databaseUrlMysql

	// call database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return
	}

	fmt.Println("Database Connected")
	// db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
}
