package config

import (
	"eduquiz-api/model/schema"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to Connect Database")
	}

	fmt.Println("Connected to Database")

}

func Migrate() {

	err := DB.AutoMigrate(&schema.Student{}, &schema.Teacher{}, &schema.QuizCategory{}, &schema.Quiz{}, &schema.Question{}, &schema.KeyAnswer{}, &schema.AttemptAnswer{}, &schema.QuizResult{})
	if err != nil {
		log.Fatal("Failed to Migrate Database")
	}
	fmt.Println("Success Migrate Database")
}
