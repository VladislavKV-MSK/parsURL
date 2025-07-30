package storage

import (
	"github.com/VladislavKV-MSK/parsURL/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(url *models.URL) error
	FindByShortCode(shortCode string) (*models.URL, error)
}

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortCode(shortCode string) (*models.URL, error) {
	var url models.URL
	err := r.db.Where("short_code = ?", shortCode).First(&url).Error
	return &url, err
}
