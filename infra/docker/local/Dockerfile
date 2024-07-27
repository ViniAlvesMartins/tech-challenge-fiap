FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.51.0
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN go build -o ./cmd/api/main ./cmd/api/main.go
RUN go build -o ./cmd/worker/cancel_awaiting_payment_orders/main ./cmd/worker/cancel_awaiting_payment_orders/main.go
RUN go build -o ./cmd/worker/update_order_status/main ./cmd/worker/update_order_status/main.go
RUN go build -o ./cmd/worker/rollback_payment/main ./cmd/worker/rollback_payment/main.go

EXPOSE 8081

CMD "air"
