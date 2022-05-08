FROM alpine
ENV VERSION 1.0
MAINTAINER yangshuangxi<749815394@qq.com>
ADD ./server /usr/local/bin/server
EXPOSE 7848
CMD ["/usr/local/bin/server"]
