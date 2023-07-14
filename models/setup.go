package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUrl := os.Getenv("DB_URL")

	DB, err = gorm.Open(dbDriver, dbUrl)

	if err != nil {
		fmt.Println("Cannot connect to database ", dbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", dbDriver)
	}

	DB.AutoMigrate(&User{})

}
