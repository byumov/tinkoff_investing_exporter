# Tinkoff investing exporter for Prometheus

[![.github/workflows/mail.yml](https://github.com/byumov/tinkoff_investing_exporter/workflows/.github/workflows/mail.yml/badge.svg)](https://github.com/byumov/tinkoff_investing_exporter/actions)
[![Docker Pulls](https://img.shields.io/docker/pulls/byumov/tinkoff_investing_exporter.svg)](https://hub.docker.com/r/byumov/tinkoff_investing_exporter)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/byumov/tinkoff_investing_exporter.svg)](https://hub.docker.com/r/byumov/tinkoff_investing_exporter/builds)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/byumov/tinkoff_investing_exporter.svg)](https://hub.docker.com/r/byumov/tinkoff_investing_exporter/builds)

## Usage

First, get token from Tinkoff: [https://tinkoffcreditsystems.github.io/invest-openapi/auth/](https://tinkoffcreditsystems.github.io/invest-openapi/auth/)

### Docker

```bash
docker run -d -p 2112:2112 -e TCS_TOKEN="YOUR_TOKEN" byumov/prometheus_tcs
curl 127.0.0.1:2112/metrics
```

### Binary

```bash
TCS_TOKEN="YOUR_TOKEN" ./tinkoff_investing_exporter
```

## Configuration via env variables

`TCS_UPDATE_INTERVAL` in seconds, 120 by default. Be careful, Tinkoff API has a limit [120 requests\min](https://tinkoffcreditsystems.github.io/invest-openapi/rest/)

`TCS_DEBUG`, set with any value for more verbosity

`TCS_LISTEN_PORT`, 2112 by default

## Example dashboard

![Example dashboard](https://i.imgur.com/ixBQmug.png)

## How to build

### Binary

```bash
git clone https://github.com/byumov/tinkoff_investing_exporter.git
cd tinkoff_investing_exporter
GOOS="linux" GOARCH="amd64" go build .
```

For other OS see `GOOS` and `GOARCH` values in [go documentation](https://golang.org/doc/install/source#environment).

### Docker image

```bash
git clone https://github.com/byumov/tinkoff_investing_exporter.git
cd tinkoff_investing_exporter
docker build .
```
