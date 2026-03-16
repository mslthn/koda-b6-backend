FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/coffee-backend ./cmd/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/coffee-backend .

# COPY --from=builder /app/.env .

COPY --from=builder /app/uploads ./uploads

RUN apk add --no-cache tzdata

RUN chmod +x /app/coffee-backend

EXPOSE 8888

ENTRYPOINT ["./coffee-backend"]