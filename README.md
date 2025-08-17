# URL Shortener Service

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8.svg?logo=go)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-1.9+-lightblue.svg)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-1.25+-FF69B4.svg)](https://gorm.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791.svg?logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-20.10+-2496ED.svg?logo=docker)](https://www.docker.com/)

[![HTTP](https://img.shields.io/badge/Protocol-HTTP-orange.svg)](https://developer.mozilla.org/en-US/docs/Web/HTTP)
[![REST](https://img.shields.io/badge/API-REST-green.svg)](https://en.wikipedia.org/wiki/Representational_state_transfer)
[![JSON](https://img.shields.io/badge/Data-JSON-yellow.svg)](https://www.json.org/)
[![HTML5](https://img.shields.io/badge/Frontend-HTML5-E34F26.svg?logo=html5)](https://developer.mozilla.org/en-US/docs/Web/HTML)
[![CSS3](https://img.shields.io/badge/Style-CSS3-1572B6.svg?logo=css3)](https://developer.mozilla.org/en-US/docs/Web/CSS)

## 🇬🇧 English
Microservice for URL shortening with web interface, using Gin, GORM and PostgreSQL.

## 🚀 Features

- REST API for integration
- Shorten long URLs to compact codes (e.g. http://example.com/abc123)
- Web interface for easy URL shortening
- Redirect from short code to original URL
- Duplicate detection (returns existing code for previously shortened URLs)
- Automatic PostgreSQL table creation
- Database connection retries
- URL validation

## 🌐 Web Interface

### Features:
- Simple URL input form
- Short URL display with copy button
- Long URL support (auto text wrapping)
- Visual duplicate indication
- Responsive design

## 💻 Technologies
- **Go 1.24** (with modules)
- **Gin** - HTTP framework
- **GORM** - PostgreSQL ORM
- **PostgreSQL 13** - database
- **Docker** - containerization
- **HTML5/CSS3** - web interface
- **JavaScript** - frontend interactivity

## 🛠 Installation

### Local Development
Service will be available at http://localhost:8000 (frontend) and http://localhost:8080 (API)

1. Clone repository

```bash
git clone https://github.com/VladislavKV-MSK/parsURL.git
cd parsURL
```
2. Run with Docker Compose

```bash
docker-compose up --build
```
3. Open in browser:

```text
http://localhost:8000
```

## ⚙️ Configuration

| Variable       | Default               | Description                  |
|----------------|-----------------------|------------------------------|
| `DB_HOST`      | `db`                  | PostgreSQL host              |
| `DB_PORT`      | `5432`                | PostgreSQL port              |
| `DB_USER`      | `postgres`            | Database user                |
| `DB_PASSWORD`  | `postgres`            | Database password            |
| `DB_NAME`      | `url_shortener`       | Database name                |
| `SERVER_PORT`  | `8080`                | HTTP server port             |
| `BASE_URL`     | `http://localhost:8080` | Base URL for short links     |

## 📡 API Endpoints

### 1. Shorten URL:

```http
POST /api/shorten
Content-Type: application/json

{
  "url": "https://example.com/very/long/url"
}
```
**Response:**

```json
{
  "short_url": "http://short.example.com/abc123",
  "is_duplicate": false,
  "created_at": "2025-03-15T14:30:00Z"
}
```
### 2. Redirect
```http
GET /{shortCode}
```

## 🐳 Docker Deployment
**Locally**
```bash
docker-compose up --build
```
**On server**
- Configure reverse proxy (Nginx/Apache)
- Add SSL certificate (Let's Encrypt)
- Run in background:
```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

## 📄 License
Project is MIT licensed:
```text
MIT License
Copyright (c) 2025 VladislavKV_MSK
```

## 🇷🇺 Русский
Микросервис для сокращения длинных URL-адресов с использованием Gin, GORM и PostgreSQL.

## 🚀 Функционал

- Сокращение длинных URL в короткие коды (например: `http://example.com/abc123`)
- Веб-интерфейс для удобного сокращения URL (будет доступен на: `[http](http://localhost:8000)`)
- Определение дубликатов (возвращает существующий код для уже сокращенных URL)
- Редирект по короткому коду на оригинальный URL
- Автоматическое создание таблиц в PostgreSQL
- Повторные попытки подключения к БД при старте
- Валидация входных URL

## 🌐 Веб-интерфейс

Особенности:
- Простая форма для ввода URL
- Отображение сокращенной ссылки с кнопкой копирования

## 💻 Технологии

- **Go 1.24** (с модулями)
- **Gin** - HTTP-фреймворк
- **GORM** - ORM для работы с PostgreSQL
- **PostgreSQL 13** - база данных
- **Docker** - контейнеризация
- **HTML5/CSS3** - веб-интерфейс
- **JavaScript** - интерактивность фронтенда

## 🛠 Запуск

Сервис будет доступен на http://localhost:8000 (фронтенд) и http://localhost:8080 (API)

### 1. Клонирование репозитория
```bash
git clone https://github.com/VladislavKV-MSK/parsURL.git
cd parsURL
```
### 2. Docker Compose (основной способ)
```bash
# Сборка и запуск (с автоматическим созданием БД)
# С дефолтными переменными
docker-compose up --build

# Остановка и удаление контейнеров
docker-compose down

# Запуск с кастомными переменными (дополнительно)
POSTGRES_USER=myadmin \
POSTGRES_PASSWORD=secret123 \
DB_USER=appuser \
DB_PASSWORD=apppass \
SERVER_PORT=9090 \
docker-compose up
```

Либо переопределяете переменные в docker-compose.yml, либо создаете конфигурационный файл .env с нужынми перменными и запускаете docker-compose up --build

## ⚙️ Конфигурация
| Переменная      | По умолчанию    | Описание                |
|----------------|----------------|-------------------------|
| `DB_HOST`      | `localhost`     | Хост PostgreSQL         |
| `DB_PORT`      | `5432`          | Порт PostgreSQL         |
| `DB_USER`      | `postgres`      | Пользователь БД         |
| `DB_PASSWORD`  | `postgres`      | Пароль БД              |
| `DB_NAME`      | `url_shortener` | Имя базы данных        |
| `SERVER_PORT`  | `8080`          | Порт HTTP-сервера      |


## 🧪 Тестирование

Запуск тестов:

```bash
go test -v ./...
```
Тесты покрывают:
- Валидацию URL
- Обработку ошибок БД
- Редиректы

## Проверка работоспособности (CURL)

### 1. Сокращение URL
```bash
curl -X POST "http://localhost:8080/shorten" \
     -H "Content-Type: application/json" \
     -d '{"url":"https://google.com"}'
```
Ответ:

```json
{
  "short_url": "http://short.example.com/abc123",
  "is_duplicate": false
}
```
### 2. Редирект по короткой ссылке
```bash
curl -v "http://localhost:8080/abc123"
```
Ожидаемый результат:
```text
HTTP/1.1 301 Moved Permanently
Location: https://google.com
```
### 3. Проверка ошибок
```bash
# Невалидный URL
curl -X POST "http://localhost:8080/shorten" \
     -d '{"url":"invalid-url"}'

# Несуществующий код
curl -v "http://localhost:8080/badcode"
```
Замените abc123 на реальный код из ответа сервера.

## 📂 Структура проекта
```text
.
├── config/               # Конфигурация
│   └── config.go
├── handlers/             # HTTP-обработчики
│   ├── url_handler.go
│   └── url_test.go
├── models/               # Модели данных
│   └── url.go
├── storage/              # Работа с БД
│   ├── db.go
│   └── url_repo.go
├── web.html              # Веб-интерфейс
├── docker-compose.yml    # Конфигурация Docker
├── Dockerfile            # Конфигурация Docker
├── Dockerfile.frontend   # Конфигурация Docker для web
├── go.mod                # Зависимости Go
└── main.go               # Точка входа
```
## 📄 Лицензия
Проект распространяется под лицензией MIT:

```text
MIT License
Copyright (c) 2025 VladislavKV_MSK
```
