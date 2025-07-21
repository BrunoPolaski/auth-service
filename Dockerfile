FROM golang:1.25-rc-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service .

FROM golang:1.25-rc-alpine AS runner

RUN adduser -D authuser

COPY --from=builder /app/auth-service /app/auth-service
COPY --from=builder /app/.env /app/.env

RUN chown -R authuser:authuser /app
RUN chmod +x /app/auth-service

WORKDIR /app

EXPOSE 8080

USER authuser

CMD ["./auth-service"]