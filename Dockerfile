FROM golang:1.14 AS build
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build .

FROM alpine:3.12
COPY --from=build /app/tinkoff_investing_exporter /app/tinkoff_investing_exporter
EXPOSE 2112
CMD /app/tinkoff_investing_exporter
