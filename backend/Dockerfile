FROM golang:1.23.0-alpine3.19

WORKDIR /src/app

RUN go install github.com/cosmtrek/air@v1.28.0

COPY . .

RUN go mod tidy
