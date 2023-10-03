package database

import (
	"github.com/szmulinho/feedback/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	connection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=L96a1prosniper dbname=feedback port=5433 sslmode=disable TimeZone=Europe/Warsaw"), &gorm.Config{})

	if err != nil {
		panic("can't connect with database")
	}

	DB = connection

	connection.AutoMigrate(&model.Opinion{})

	return connection
}