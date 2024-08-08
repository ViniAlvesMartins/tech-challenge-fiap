FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

EXPOSE 8080

RUN go build -o ./cmd/api/main ./cmd/api/main.go
RUN go build -o ./cmd/worker/cancel_awaiting_payment_orders/main ./cmd/worker/cancel_awaiting_payment_orders/main.go
RUN go build -o ./cmd/worker/update_order_status/main ./cmd/worker/update_order_status/main.go
RUN go build -o ./cmd/worker/update_payment_status/main ./cmd/worker/update_payment_status/main.go

CMD [ "./cmd/api/main" ]