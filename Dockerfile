FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /root

# Копируем бинарник из builder
COPY --from=builder /app/app .

# Создаём папку для логов (если используешь файловый вывод)
RUN mkdir -p /logs

# Переменные окружения (если нужно)
ENV LOG_FILE=/logs/app.log

# Запуск приложения
CMD ["./app"]