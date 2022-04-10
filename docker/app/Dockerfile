FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN go mod init app/main

RUN go get github.com/gin-gonic/gin \
    github.com/go-sql-driver/mysql \
    github.com/jinzhu/gorm
