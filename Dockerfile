# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.21

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/main .
COPY --from=builder  /app/migrate ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x /app/start.sh
COPY db/migrate ./migration

EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
