FROM golang:1.14-alpine3.12
MAINTAINER bayu3490

ENV GIN_MODE=release

WORKDIR /app
COPY ./ /app
RUN go build -o main.out main.go
CMD ["./main.go"]