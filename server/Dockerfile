#билд бинарника
FROM golang:latest AS builder
WORKDIR /app
COPY .. .
RUN go build -o main main.go
#запуск
FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 8080
CMD  ["/app/main"]
