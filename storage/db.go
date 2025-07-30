package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/VladislavKV-MSK/parsURL/config"
	"github.com/VladislavKV-MSK/parsURL/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Пытаемся подключиться 5 раз с задержкой
	for i := 0; i < 5; i++ {
		db, err = NewDB(cfg)
		if err == nil {
			return db, nil
		}
		log.Printf("Failed to connect to DB (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	return nil, err
}

func NewDB(cfg *config.Config) (*gorm.DB, error) {

	// Подключение к БД, которая создается через Docker
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), models.GormConfig)
	if err != nil {
		return nil, err
	}

	// Автомиграция (создаст таблицу, если её нет)
	if err := db.AutoMigrate(&models.URL{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return db, nil
}
