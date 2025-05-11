FROM golang:1.24.2-alpine

EXPOSE 8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .env ./

COPY . .

RUN go build -o main cmd/main.go

CMD ["./main"]