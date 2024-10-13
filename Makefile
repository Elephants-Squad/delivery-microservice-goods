# Установи переменные для подключения к БД
DB_DSN := "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(POSTGRES_DB)?sslmode=disable"

# Команда для запуска миграций
migrate-up:
	go run backend/cmd/main.go migrate-up

# Команда для отката миграций
migrate-down:
	go run cmd/main.go migrate-down

# Запуск сервера
run:
	go run backend/cmd/main.go

# Тестирование
test:
	go test ./...

# Билд
build:
	go build -o bin/app cmd/main.go
