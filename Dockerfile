# Указываем базовый образ с поддержкой Go
FROM golang:alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Создаем исполняемую и служебную директории
RUN mkdir cmd && mkdir internal

# Копируем в них файлы
COPY ./cmd ./cmd
COPY ./internal ./internal

# Устанавливаем зависимости
RUN go mod download

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

# Устанавливаем свои переменные окружения для работы приложения
ENV BOT_TOKEN=""
ENV API_KEY=""

# Указываем команду запуска
CMD ["./main"]
