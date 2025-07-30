# URL Shortener Service

Микросервис для сокращения длинных URL-адресов с использованием Gin, GORM и PostgreSQL.

## 🚀 Функционал

- Сокращение длинных URL в короткие коды (например: `http://example.com/abc123`)
- Редирект по короткому коду на оригинальный URL
- Автоматическое создание таблиц в PostgreSQL
- Повторные попытки подключения к БД при старте
- Валидация входных URL

## 💻 Технологии

- **Go 1.24** (с модулями)
- **Gin** - HTTP-фреймворк
- **GORM** - ORM для работы с PostgreSQL
- **PostgreSQL 13** - база данных
- **Docker** - контейнеризация

## 🛠 Запуск

Сервис будет доступен на http://localhost:8080

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

# Запуск с кастомными переменными
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
{"short_url":"http://localhost:8080/abc123"}
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
├── docker-compose.yml    # Конфигурация Docker
├── Dockerfile            # Конфигурация Docker
├── go.mod                # Зависимости Go
└── main.go               # Точка входа
```
## 📄 Лицензия
Проект распространяется под лицензией MIT:

```text
MIT License

Copyright (c) 2025 VladislavKV_MSK
```
