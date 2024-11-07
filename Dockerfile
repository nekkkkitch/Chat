FROM golang:1.23-alpine3.20
WORKDIR /app
COPY go.mod go.sum cfg.yml ./
RUN go mod download