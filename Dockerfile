FROM golang:1.20-bullseye as golang
LABEL maintainer="i@chengxuan.li"

ENV GOPROXY=https://goproxy.cn,direct \
GO111MODULE=auto
ADD ./ /gocode
RUN cd /gocode && \
    go build ./


FROM alpine:3.18.2
COPY --from=golang /gocode/po2json /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/po2json"]
