FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum cfg.yml ./
RUN go mod download