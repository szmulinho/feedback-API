package database

import (
	"github.com/szmulinho/feedback/internal/config"
	"github.com/szmulinho/feedback/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conn := config.LoadConfigFromEnv()
	connectionString := conn.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Opinion{}); err != nil {
		return nil, err
	}

	return db, nil
}
