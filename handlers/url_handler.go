package handlers

import (
	"errors"
	"net/http"

	"github.com/VladislavKV-MSK/parsURL/models"
	"github.com/VladislavKV-MSK/parsURL/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// Обработчик для POST /shorten
func ShortenURL(repo storage.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Парсим входной JSON
		var req ShortenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			// Если JSON невалиден — возвращаем ошибку 400
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Сначала проверяем, есть ли уже такой URL в БД
		existingURL, err := repo.FindByOriginalURL(req.URL)
		if err == nil {
			// URL уже существует - возвращаем существующую короткую ссылку
			scheme := "http://"
			if c.Request.TLS != nil {
				scheme = "https://"
			}
			c.JSON(http.StatusOK, gin.H{
				"short_url":    scheme + c.Request.Host + "/" + existingURL.ShortCode,
				"is_duplicate": true, // Добавляем флаг, что это существующая запись
			})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) { // Если ошибка и это не "не найдено" - возвращаем ошибку
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// Создаем короткий URL код
		shortCode := uuid.New().String()[:8]

		// Создаём запись в БД
		url := &models.URL{
			OriginalURL: req.URL,   // Оригинальный URL из запроса
			ShortCode:   shortCode, // Сгенерированный код
		}

		if err := repo.Create(url); err != nil {
			// Если ошибка при сохранении — возвращаем 500
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create short URL"})
			return
		}

		scheme := "http://"
		if c.Request.TLS != nil {
			scheme = "https://"
		}
		c.JSON(http.StatusOK, gin.H{
			"short_url":    scheme + c.Request.Host + "/" + shortCode, // Полная короткая ссылка
			"is_duplicate": false,
		})
	}
}

// Обработчик для GET /:shortCode
func RedirectURL(repo storage.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем shortCode из URL
		shortCode := c.Param("shortCode")

		// Ищем URL в БД по коду
		url, err := repo.FindByShortCode(shortCode)
		if err != nil {
			// Если не найдено — возвращаем 404
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		// Делаем редирект (301 — "Moved Permanently")
		c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
	}
}
