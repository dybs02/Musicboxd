FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env.sample .
EXPOSE 8080
CMD ["./main"]
