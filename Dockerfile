FROM golang:1.17.2-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go-jokes-api

EXPOSE $PORT

CMD "./go-jokes-api"
