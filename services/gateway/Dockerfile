FROM golang:1.23
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o router ./services/gateway/cmd/app/main.go