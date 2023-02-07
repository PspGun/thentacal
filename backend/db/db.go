package db

import (
	"github.com/PspGun/thentacal/type/database"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBsetup() (err error) {

	dsn := os.Getenv("DB_ENV")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	db.AutoMigrate(&database.User{})
	db.AutoMigrate(&database.Report{})
	return err
}
