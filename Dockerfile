# Указываем версию Go и базовый образ (alpine)
ARG GO_VERSION=1.22.5
ARG GO_BASE_IMAGE=alpine

# Используем базовый образ Go
FROM golang:${GO_VERSION}-${GO_BASE_IMAGE} AS builder

WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY go.mod go.sum ./

RUN go mod download

# Копируем содержимое папки backend
COPY . .

# Сборка приложения
RUN go build -o ./backend/main ./backend/cmd/main.go
