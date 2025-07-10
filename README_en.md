# http-health-probe

[![Go Report Card](https://goreportcard.com/badge/github.com/dev-shimada/http-health-probe)](https://goreportcard.com/report/github.com/dev-shimada/http-health-probe)
[![CI](https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml/badge.svg)](https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml)
[![Coverage Status](https://coveralls.io/repos/github/dev-shimada/http-health-probe/badge.svg?branch=main)](https://coveralls.io/github/dev-shimada/http-health-probe?branch=main)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/dev-shimada/http-health-probe/blob/master/LICENSE)

http-health-probe is a simple command-line tool for health checking HTTP-based APIs. 

## ✨ Features

- No C language libraries required (implemented in Go only)
- Works with distroless images
- Simple command-line interface

## 📦 Installation

### Download Binary

Download the binary for your platform from the releases page.
```shell
INSTALL_DIR=/tmp
VERSION=v0.1.1
TARGETOS=Linux
TARGETARCH=arm64
curl --silent --location https://github.com/dev-shimada/http-health-probe/releases/download/${VERSION}/http-health-probe_${TARGETOS}_${TARGETARCH}.tar.gz | tar xvz -C ${INSTALL_DIR} --one-top-level=http-health-probe_${TARGETOS}_${TARGETARCH}
```

### go install

```shell
go install github.com/dev-shimada/http-health-probe
```

## 🚀 Usage

### Command Example

```shell
http-health-probe --addr=http://localhost:8080/healthcheck
```

### ⚙️ Main Options

| Option              | Description                                             | Default Value       |
| ------------------- | ------------------------------------------------------- | ------------------- |
| `--addr`            | HTTP endpoint URL to check                              | None (required)     |
| `--method`          | HTTP method                                             | `GET`               |
| `--timeout`         | Request timeout                                         | `1s`                |
| `--expected-status` | Expected HTTP status code for a successful health check | `200`               |
| `--tls`             | Use TLS                                                 | `false`             |
| `--insecure`        | Skip TLS certificate verification                       | `false`             |
| `--user-agent`      | User-Agent header                                       | `http-health-probe` |

## 🐳 Example with distroless image

```Dockerfile
FROM rust:bookworm AS build
ARG TARGETOS
ARG TARGETARCH
RUN <<EOF
INSTALL_DIR=/tmp
VERSION=v0.1.0
curl --silent --location https://github.com/dev-shimada/http-health-probe/releases/download/${VERSION}/http-health-probe_${TARGETOS}_${TARGETARCH}.tar.gz | tar xvz -C ${INSTALL_DIR} --one-top-level=http-health-probe_${TARGETOS}_${TARGETARCH}
EOF

FROM gcr.io/distroless/base-debian12:latest
ARG TARGETOS
ARG TARGETARCH
COPY --chown=nonroot:nonroot --from=build /tmp/http-health-probe_${TARGETOS}_${TARGETARCH}/http-health-probe /bin/http-health-probe
HEALTHCHECK --interval=10s --timeout=3s --start-period=5s CMD ["/bin/http-health-probe", "--addr=:3000"]
USER nonroot:nonroot
```

## 📝 License

This project is licensed under the MIT License.

## 🤝 Contribution

Issues and Pull Requests are welcome. Feel free to report bugs or suggest features.
