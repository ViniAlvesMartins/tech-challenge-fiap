FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENV GOPATH=/go

RUN CGO_ENABLED=0 GOOS=linux go build -o /tech-challenge-fiap

CMD ["tech-challenge-fiap"]