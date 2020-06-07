# Tinkoff investing exporter

## Usage

First, get token from Tinkoff: [https://tinkoffcreditsystems.github.io/invest-openapi/auth/](https://tinkoffcreditsystems.github.io/invest-openapi/auth/)

Docker:

```bash
docker run -d -p 2112:2112 -e TCS_TOKEN="YOUR_TOKEN" byumov/prometheus_tcs
curl 127.0.0.1:2112/metrics
```

## How to build

Binary:

```bash
git clone https://github.com/byumov/tinkoff_investing_exporter.git
cd tinkoff_investing_exporter
GOOS="linux" GOARCH="amd64" go build .
```

For other OS see `GOOS` and `GOARCH` values in [go documentation](https://golang.org/doc/install/source#environment).

Docker image:

```bash
git clone https://github.com/byumov/tinkoff_investing_exporter.git
cd tinkoff_investing_exporter
docker build .
```
