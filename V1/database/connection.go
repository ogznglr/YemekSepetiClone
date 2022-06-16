package database

import (
	"fmt"
	"yemeksepeti/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "superuser"
	dbname   = "yemeksepeti_db"
)

func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("No connection to the DataBase")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Order{})
	DB = db
}
