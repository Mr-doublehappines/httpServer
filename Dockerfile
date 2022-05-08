FROM golang:1.17 AS builder
WORKDIR /httpserver
COPY ./server.go /httpserver/
ENV GO111MODULE=off \
    CGO_ENABLED=0    \
    GOARCH=amd64    \
    GOOS=linux
RUN go build  -o server ./server.go 

FROM alpine
ENV VERSION 1.0
MAINTAINER yangshuangxi<749815394@qq.com>
COPY --from=builder /httpserver/server  /usr/local/bin/
EXPOSE 7848
CMD ["/usr/local/bin/server"]
