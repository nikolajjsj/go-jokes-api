FROM golang:1.17.2-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/go-jokes-api /app/
WORKDIR /app
CMD ["./go-jokes-api"]
