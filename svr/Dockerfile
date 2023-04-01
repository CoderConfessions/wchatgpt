FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk update && \
    apk upgrade && \ 
    apk add --no-cache gcc musl-dev

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" ./main.go

EXPOSE 8080

ENTRYPOINT ["/app/main"]