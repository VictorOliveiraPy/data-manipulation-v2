FROM golang:1.20-alpine

ADD . /app

WORKDIR /app
ENV  go -w GO111MODULE=off


RUN go get -d -v ./...
RUN go install -v ./...
