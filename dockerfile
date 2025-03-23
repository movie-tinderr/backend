# Используем официальный образ Go для сборки
FROM golang:1.24.1-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum (если они есть)
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем приложение
RUN go build -o my-go-app .

# Используем минимальный образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из builder
COPY --from=builder /app/my-go-app /app/my-go-app

# Указываем команду для запуска приложения
CMD ["/app/my-go-app"]