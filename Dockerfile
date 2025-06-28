FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .env ./

COPY . .

RUN go build cmd/main.go

CMD ["./main"]