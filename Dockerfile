FROM golang:latest

LABEL maintainer="gRPC server"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN apt-get update && apt-get install make

COPY . .

RUN go build -o main ./cmd

CMD make run