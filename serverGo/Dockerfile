#билд бинарника
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY .. .

RUN go build -o main main.go

#запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY считать.xlsx .
COPY db/migrations ./db/migrations

EXPOSE 8080
CMD  ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
