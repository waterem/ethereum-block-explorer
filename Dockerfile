FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

ENV GOBIN /go/bin

RUN go get github.com/urfave/cli
RUN go get github.com/go-sql-driver/mysql


