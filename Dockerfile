# Use the official Golang image as a build stage
FROM golang:1.22.1

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/mynance-service-api/main.go

EXPOSE 8080

CMD [ "/app/main" ]