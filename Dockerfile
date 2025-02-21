# syntax=docker/dockerfile:1
# A basic app to experiment with go
FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-demo-app

EXPOSE 8080

CMD ["/go-demo-app"]