FROM golang:1.15
WORKDIR /go/src/app

RUN go get github.com/go-sql-driver/mysql
# mysqlにつなぐために必要
RUN go get github.com/jinzhu/gor
# db接続に必要