FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
RUN adduser -D appuser
USER appuser
WORKDIR /home/appuser
COPY --from=builder --chown=appuser:appuser /app/main .
CMD ["./main"]
