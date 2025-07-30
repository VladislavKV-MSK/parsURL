package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type URL struct {
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"unique;not null"`
}

var GormConfig = &gorm.Config{
	// Отключаем логирование (включить для дебага)
	Logger: logger.Default.LogMode(logger.Silent),
}
