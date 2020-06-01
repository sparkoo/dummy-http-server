FROM golang:1.14 AS builder
WORKDIR /go/src
COPY main.go .
COPY templates /go/bin/templates
RUN go build -o /go/bin/server main.go
WORKDIR /go/bin
CMD ["/go/bin/server"]
