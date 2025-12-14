package postgres

import (
	"fmt"
	"log"
	"task-manager/internal/config"
	"task-manager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.SSLMode,
	)

	if cfg.DBPassword != "" {
		dsn += fmt.Sprintf(" password=%s", cfg.DBPassword)
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db.AutoMigrate(&models.Task{})

	return db
}
