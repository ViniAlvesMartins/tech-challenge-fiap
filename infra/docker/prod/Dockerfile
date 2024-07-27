FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

EXPOSE 8081

RUN go build -o ./cmd/api/main ./cmd/api/main.go
RUN go build -o ./cmd/worker/cancel_awaiting_payment_orders/main ./cmd/worker/cancel_awaiting_payment_orders/main.go
RUN go build -o ./cmd/worker/update_order_status/main ./cmd/worker/update_order_status/main.go
RUN go build -o ./cmd/worker/rollback_payment/main ./cmd/worker/rollback_payment/main.go

CMD [ "./cmd/api/main" ]