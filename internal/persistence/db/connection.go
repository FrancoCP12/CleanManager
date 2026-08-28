package db

import (
	"fmt"
	"log"
	"time"

	"github.com/francocp12/CleanManager/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var err error
	var dsl string

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	dsl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsl), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting native sqlDB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return db
}
