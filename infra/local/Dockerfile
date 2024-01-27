FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD "air"
