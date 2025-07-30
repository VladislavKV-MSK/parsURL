package main

import (
	"log"

	"github.com/VladislavKV-MSK/parsURL/config"
	"github.com/VladislavKV-MSK/parsURL/handlers"
	"github.com/VladislavKV-MSK/parsURL/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	// Загружаем конфиг
	cfg := config.Load()

	// Инициализация БД
	db, err := storage.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Could not connect to DB after retries:", err)
	}
	repo := storage.NewURLRepository(db)

	// Инициализация Gin
	r := gin.Default()

	// Инициализация handlers

	r.POST("/shorten", handlers.ShortenURL(repo))
	r.GET("/:shortCode", handlers.RedirectURL(repo))

	// Запуск сервера
	r.Run(":8080")
}
