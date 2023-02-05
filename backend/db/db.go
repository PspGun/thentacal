package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}


func DBsetup() (err error) {
	dsn := os.Getenv("DB_ENV")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	db.AutoMigrate(&User{})
	return err
}