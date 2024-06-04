# Используем базовый образ golang:1.21-alpine3.18
FROM golang:1.21-alpine3.18 AS builder
# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app
# Копируем файлы из текущего контекста сборки (main.go)
COPY . .
# Собираем проект
RUN go build -o main main.go
# Используем базовый образ alpine:latest
FROM alpine:latest
# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app
# Копируем исполняемый файл и другие необходимые файлы из сборщика
COPY --from=builder /app/main /app/
COPY --from=builder /app/app.env /app/
COPY --from=builder /app/start.sh /app/
COPY --from=builder /app/wait-for.sh /app/
COPY --from=builder /app/считать.xlsx /app/
COPY --from=builder /app/db/migrations /app/db/migrations
# Открываем порт 8080
EXPOSE 8080
# Команда запуска приложения
CMD ["./start.sh"]
# Команда для входа при запуске контейнера
ENTRYPOINT ["sh", "-c"]
