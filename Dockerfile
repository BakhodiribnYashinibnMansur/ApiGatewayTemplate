FROM golang:1.19.1-alpine3.16

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go build -o main ./command/main.go
RUN go mod tidy
RUN go mod vendor

CMD ["./main"]
