FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url-shortener .

RUN apk add --no-cache ca-certificates

CMD ["./url-shortener"]