FROM golang:1.21

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build -o /server

CMD [ "/server" ]
