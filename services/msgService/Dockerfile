FROM nekkkkitch/docker AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o msg ./services/msgService/cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/msgService .
COPY ./services/msgService/cfg.yml .
CMD ["./msgService"]