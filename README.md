# http-health-probe

<p align="center">
  <a href="https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml"><img alt="CI" src="https://img.shields.io/github/actions/workflow/status/dev-shimada/http-health-probe/CI.yaml?branch=main&style=flat-square&logo=githubactions&logoColor=white&label=CI"></a>
  <a href="https://coveralls.io/github/dev-shimada/http-health-probe?branch=main"><img alt="Coverage" src="https://img.shields.io/coveralls/github/dev-shimada/http-health-probe/main?style=flat-square&logo=coveralls&logoColor=white"></a>
  <a href="https://github.com/dev-shimada/http-health-probe/blob/main/go.mod"><img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/dev-shimada/http-health-probe?style=flat-square&logo=go&logoColor=white"></a>
  <a href="https://github.com/dev-shimada/http-health-probe/blob/main/LICENSE"><img alt="MIT License" src="https://img.shields.io/github/license/dev-shimada/http-health-probe?style=flat-square"></a>
</p>

<p align="center">
  <a href="https://github.com/dev-shimada/http-health-probe/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/dev-shimada/http-health-probe?style=flat-square&logo=github&logoColor=white"></a>
  <a href="#-installation"><img alt="go install" src="https://img.shields.io/badge/go%20install-dev--shimada%2Fhttp--health--probe-00ADD8?style=flat-square&logo=go&logoColor=white"></a>
</p>

http-health-probe is a simple command-line tool for health checking HTTP-based APIs. 

*(日本語版は[こちら](./README.ja.md))*

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
VERSION=v0.1.1
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
