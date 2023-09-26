FROM golang:1.21 AS build

WORKDIR /app

COPY . .

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]

# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /server /server

# EXPOSE 8080

# USER nonroot:nonroot

# CMD ["/server"]
