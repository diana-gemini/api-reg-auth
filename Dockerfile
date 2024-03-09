FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && go build -o api ./cmd/*
FROM alpine:3.6
LABEL Authors="@dtyuligu" Project="API"
WORKDIR /app
COPY --from=builder /app .
CMD ["/app/api"]