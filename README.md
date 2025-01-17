# Лабораторный API

Проект представляет собой простейший пример описание API и интеграции с базой данных PostgreSQL. 
Используются структурированные подходы для тестирования, обработки конфигураций и документации. 

API построен на Go с использованием `Gin`, `GoDoc` и `Goose` для миграций базы данных.

## Структура проекта

Проект организован следующим образом:
- `bin/` - скомпилированные файлы
- `cmd/` - основная директория с `main.go` для запуска сервера.
- `internal/handler/` - содержит логику обработки HTTP-запросов.
- `internal/repository/` - содержит реализации доступа к базе данных.
- `internal/routes/` - содержит генерацию маршрутов API
- `internal/config/` - модуль конфигурации для загрузки переменных из `.env` файла.
- `db/migrations/` - содержит SQL-миграции для базы данных, управляемые через Goose.
- `db/queries/` - содержатся запросы к БД в формате .sql
- `tests/` - директория для тестов API и интеграционных тестов.

## Зависимости

- Go 1.23.2
- Gin
- Goose (для управления миграциями)
- PostgreSQL
- Godoc
- Swagger

## Установка и настройка

### Шаг 1: Клонирование репозитория

```bash
git clone https://github.com/kyvees13/lab-server.git
cd lab-server
```

```bash
# для загрузки зависимостей
make deps
```

### Шаг 2: Конфигурация переменных окружения
Следует скопировать файл с примером `.env.example` в `.env` откуда проект считывает все переменные

```bash
cp .env.example .env
```

И заполнить его переменными по описанному шаблону

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=lab_database
```

### Шаг 3: Тестирование
Тесты находятся в папке `/tests` и используют стандартную библиотеку testing. Каждый тест проверяет конечные точки API и ожидает корректный ответ от сервера.
```bash
make test
```

### Шаг 4: Документация API
Документация API создается с использованием Swagger.

```bash
make swagger-init
```

Документация будет доступна по адресу `/swagger/index.html` после запуска сервера.



