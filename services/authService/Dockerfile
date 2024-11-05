FROM golang:1.23
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o router ./services/authService/cmd/app/main.go