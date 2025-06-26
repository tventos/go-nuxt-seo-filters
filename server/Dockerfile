FROM golang:1.22 AS builder

WORKDIR /app

ENV GOTOOLCHAIN=auto

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
