FROM golang:1.14-alpine3.12 as build-go-deps

WORKDIR /app
COPY ./ /app
RUN go build -o main.out main.go


FROM alpine:3.12.0
MAINTAINER bayu3490

RUN apk add tzdata
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=build-go-deps /app/main.out /app
ENV GIN_MODE=release
EXPOSE 8080
CMD ["./main.out"]