FROM golang:1.14.4-alpine3.12
RUN apk update && apk upgrade && \
    apk add --no-cache git
RUN go get github.com/byumov/tinkoff_investing_exporter
RUN go install github.com/byumov/tinkoff_investing_exporter
EXPOSE 2112
CMD $GOPATH/bin/tinkoff_investing_exporter
