FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o vk-bot-welcome cmd/main.go
 
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/vk-bot-welcome /app/vk-bot-welcome
COPY --from=builder /app/config /app/config

WORKDIR /app

CMD ["./vk-bot-welcome"]
