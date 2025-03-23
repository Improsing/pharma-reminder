FROM golang:1.21.3 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем Go-приложение
RUN go build -o medscheduler

# Минимальный образ для запуска
FROM debian:bullseye-slim

WORKDIR /root/

# Копируем бинарник из builder-образа
COPY --from=builder /app/medscheduler .

# Указываем порт, который слушает сервис
EXPOSE 8080

# Запускаем сервер
CMD ["./medscheduler"]
