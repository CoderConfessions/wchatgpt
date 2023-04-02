FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk update && \
    apk upgrade && \ 
    apk add --no-cache gcc musl-dev

RUN go mod download

# CGO_ENABLE=0 disable C Go support in program
# GOOS=linux set Go program built in linux
RUN CGO_ENABLED=0 GOOS=linux go build -o build/main -ldflags="-s -w" ./main.go

EXPOSE 8080

ENTRYPOINT ["/app/build/main"]