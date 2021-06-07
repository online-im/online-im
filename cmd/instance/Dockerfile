FROM golang:latest as build
ENV GO111MODULE on
ADD online-im-instance /
ADD config/glory.yaml /config/glory.yaml
ENV GLORY_CONFIG_PATH "/config/glory.yaml"
ENTRYPOINT exec /online-im-instance