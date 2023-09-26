FROM golang:1.21 AS build

WORKDIR /go/src

COPY ../../ .

RUN go install github.com/cosmtrek/air@latest

CMD "air"
