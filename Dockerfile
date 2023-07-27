FROM golang:1.20-bullseye as golang

ENV GOPROXY=https://goproxy.cn,direct \
GO111MODULE=auto
ADD ./ /gocode
RUN cd /gocode && \
    go build ./


FROM alpine:3.18.2
LABEL maintainer="i@chengxuan.li"
COPY --from=golang /gocode/po2json /usr/local/bin/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add gettext

