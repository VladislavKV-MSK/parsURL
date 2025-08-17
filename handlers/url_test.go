package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VladislavKV-MSK/parsURL/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockRepository реализует интерфейс storage.Repository для тестов
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(url *models.URL) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *MockRepository) FindByShortCode(shortCode string) (*models.URL, error) {
	args := m.Called(shortCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

// Заглушка для FindByOriginalURL - всегда возвращает "не найдено"
func (m *MockRepository) FindByOriginalURL(originalURL string) (*models.URL, error) {
	return nil, gorm.ErrRecordNotFound
}

func TestShortenURL(t *testing.T) {
	// Тестовые случаи
	tests := []struct {
		name           string
		requestBody    string
		mockSetup      func(*MockRepository)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success - valid URL",
			requestBody: `{"url":"https://example.com"}`,
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.AnythingOfType("*models.URL")).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"short_url":"http:///.*"}`,
		},
		{
			name:           "Fail - invalid URL format",
			requestBody:    `{"url":"invalid-url"}`,
			mockSetup:      func(m *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":".*"}`,
		},
		{
			name:        "Fail - repository error",
			requestBody: `{"url":"https://example.com"}`,
			mockSetup: func(m *MockRepository) {
				m.On("Create", mock.AnythingOfType("*models.URL")).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to create short URL"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок репозитория
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			// Настраиваем Gin в тестовом режиме
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.POST("/shorten", ShortenURL(mockRepo))

			// Создаем тестовый запрос
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/shorten", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Выполняем запрос
			router.ServeHTTP(w, req)

			// Проверяем результаты
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Regexp(t, tt.expectedBody, w.Body.String())

			// Проверяем, что все ожидания по моку выполнены
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRedirectURL(t *testing.T) {
	// Тестовые случаи
	tests := []struct {
		name           string
		shortCode      string
		mockSetup      func(*MockRepository)
		expectedStatus int
		expectedHeader string
	}{
		{
			name:      "Success - found URL",
			shortCode: "abc123",
			mockSetup: func(m *MockRepository) {
				m.On("FindByShortCode", "abc123").Return(&models.URL{
					OriginalURL: "https://example.com",
					ShortCode:   "abc123",
				}, nil)
			},
			expectedStatus: http.StatusMovedPermanently,
			expectedHeader: "https://example.com",
		},
		{
			name:      "Fail - not found",
			shortCode: "notfound",
			mockSetup: func(m *MockRepository) {
				m.On("FindByShortCode", "notfound").Return(nil, gorm.ErrRecordNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedHeader: "",
		},
		{
			name:      "Fail - repository error",
			shortCode: "dberror",
			mockSetup: func(m *MockRepository) {
				m.On("FindByShortCode", "dberror").Return(nil, assert.AnError)
			},
			expectedStatus: http.StatusNotFound,
			expectedHeader: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок репозитория
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			// Настраиваем Gin в тестовом режиме
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.GET("/:shortCode", RedirectURL(mockRepo))

			// Создаем тестовый запрос
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/"+tt.shortCode, nil)

			// Выполняем запрос
			router.ServeHTTP(w, req)

			// Проверяем результаты
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedHeader != "" {
				assert.Equal(t, tt.expectedHeader, w.Header().Get("Location"))
			}

			// Проверяем, что все ожидания по моку выполнены
			mockRepo.AssertExpectations(t)
		})
	}
}
