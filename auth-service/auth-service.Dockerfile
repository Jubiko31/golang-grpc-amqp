FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authService ./cmd/api

RUN chmod +x /app/authService

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authService /app

CMD ["/app/authService"]