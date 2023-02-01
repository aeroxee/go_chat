package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)
import _ "github.com/lib/pq"

var db *gorm.DB

func init() {
	db = GetDB()
	db.AutoMigrate(&User{}, &Group{}, &Chat{})
}

func GetDB() *gorm.DB {
	dsn := "host=localhost port=5432 user=fajhri password=root dbname=go_chat sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
