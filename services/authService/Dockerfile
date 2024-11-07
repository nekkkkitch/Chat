FROM nekkkkitch/docker AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o authService ./services/authService/cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/authService .
COPY ./cfg.yml .
CMD ["./authService"]