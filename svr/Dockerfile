# go build env
FROM golang:alpine AS builder
WORKDIR /src

COPY . .

ENV GOPROXY="https://goproxy.cn,direct"
RUN go mod download

# CGO_ENABLE=0 disable C Go support in program
# GOOS=linux set Go program built in linux
RUN CGO_ENABLED=0 GOOS=linux go build -o build/main -ldflags="-s -w" ./main.go

# run env
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /src/build/main .
COPY ./configuration.json /app/configuration.json

EXPOSE 8080

CMD ["/app/main", "--config-file=/app/configuration.json", "-v=4"]