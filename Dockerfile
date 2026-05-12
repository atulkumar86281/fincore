# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Run stage
FROM alpine:3.21

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD ["/app/main"]
