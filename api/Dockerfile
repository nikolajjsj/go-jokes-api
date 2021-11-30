FROM golang:1.17.2-alpine AS builder

WORKDIR /app
ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go-jokes-api

EXPOSE 8080

ENTRYPOINT ["/go-jokes-api"]