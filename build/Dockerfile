FROM golang:1.16.8-buster
MAINTAINER Jeehoon Yoo <jhyoo1220@gmail.com>

ENV DEBIAN_FRONTED noninteractive
ENV TERM linux

COPY . $GOPATH/src/github.com/jhyoo1220/go-gin-server/
WORKDIR $GOPATH/src/github.com/jhyoo1220/go-gin-server/

RUN go build -a -ldflags '-s -w' cmd/main.go

EXPOSE 8080
