package main

import (
	"log"
	"net/http"
	"time"

	"github.com/VladislavKV-MSK/parsURL/config"
	"github.com/VladislavKV-MSK/parsURL/handlers"
	"github.com/VladislavKV-MSK/parsURL/storage"

	"github.com/gin-contrib/cors"
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

	// Конфигурация CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.OPTIONS("/shorten", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8000")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusNoContent)
	})

	// Инициализация handlers

	r.POST("/shorten", handlers.ShortenURL(repo))
	r.GET("/:shortCode", handlers.RedirectURL(repo))

	// Запуск сервера
	r.Run(":8080")
}
