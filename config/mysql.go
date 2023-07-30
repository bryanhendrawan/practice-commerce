package config

import (
	"fmt"
	"log"
	"os"

	"practice-commerce/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	godotenv.Load()
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database MySQL")
	}

	log.Println("Connected to MYSQL Database")

	AutoMigrate(db)

	return db
}

func AutoMigrate(connection *gorm.DB) {
	// Notes: remove .Debug() to hide query migration
	connection.Debug().AutoMigrate(
		&entity.Merchant{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderDetail{},
		&entity.Cart{},
	)
}
