include .env

# Установите имя вашего приложения
APP_NAME = lab-server

# Генерация переменной для подключения
DB_CONN="host=$(DB_HOST) \
	port=$(DB_PORT) \
    user=$(DB_USER) \
    password=$(DB_PASS) \
    dbname=$(DB_NAME) \
    sslmode=require"

# Путь к исходным файлам и тестам
SRC_DIR = .
TEST_DIR = ./tests/...

# Команды
.PHONY: all build test clean

# Основная цель
all: build

# Сборка приложения
build:
	go build -o ./bin/ ./cmd/main.go

# Запуск тестов
test:
	go test -v $(TEST_DIR)

# Установка зависимостей
deps:
	go mod tidy
	go get -u

# Очистка скомпилированных файлов
clean:
	rm -rf ./bin/*

swagger-init:
	swag init -g cmd/main.go -dir ./ --parseInternal

# Создание новой миграции
migrate-create:
	goose -dir db/migrations create $(name) sql

migrate-status:
	goose -dir db/migrations postgres $(DB_CONN) status

# Применение всех миграций
migrate-up:
	goose -dir db/migrations postgres $(DB_CONN) up

# Откат последней миграции
migrate-down:
	goose -dir db/migrations postgres $(DB_CONN) down

# Полная команда для запуска тестов и установки зависимостей
full-test: deps test
