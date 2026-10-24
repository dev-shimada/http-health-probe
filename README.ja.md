# http-health-probe

<p align="center">
  <a href="https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml"><img alt="CI" src="https://img.shields.io/github/actions/workflow/status/dev-shimada/http-health-probe/CI.yaml?branch=main&style=flat-square&logo=githubactions&logoColor=white&label=CI"></a>
  <a href="https://coveralls.io/github/dev-shimada/http-health-probe?branch=main"><img alt="Coverage" src="https://img.shields.io/coveralls/github/dev-shimada/http-health-probe/main?style=flat-square&logo=coveralls&logoColor=white"></a>
  <a href="https://github.com/dev-shimada/http-health-probe/blob/main/go.mod"><img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/dev-shimada/http-health-probe?style=flat-square&logo=go&logoColor=white"></a>
  <a href="https://github.com/dev-shimada/http-health-probe/blob/main/LICENSE"><img alt="MIT License" src="https://img.shields.io/github/license/dev-shimada/http-health-probe?style=flat-square"></a>
</p>

<p align="center">
  <a href="https://github.com/dev-shimada/http-health-probe/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/dev-shimada/http-health-probe?style=flat-square&logo=github&logoColor=white"></a>
  <a href="https://goreportcard.com/report/github.com/dev-shimada/http-health-probe"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/dev-shimada/http-health-probe?style=flat-square"></a>
  <a href="#-インストール"><img alt="go install" src="https://img.shields.io/badge/go%20install-dev--shimada%2Fhttp--health--probe-00ADD8?style=flat-square&logo=go&logoColor=white"></a>
</p>

http-health-probeは、HTTPベースのAPIのヘルスチェックを行うためのシンプルなコマンドラインツールです。

*(英語版は[こちら](./README.md))*

## ✨ 特徴

- C言語のライブラリ不要（Goのみで実装）
- distrolessイメージで動作可能
- シンプルなコマンドラインインターフェース

## 📦 インストール

### バイナリダウンロード

リリースページから各プラットフォーム向けのバイナリをダウンロードしてください。
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

## 🚀 使い方

### コマンド例

```shell
http-health-probe --addr=http://localhost:8080/healthcheck
```

### ⚙️ 主要オプション

| オプション          | 説明                                | デフォルト値        |
| ------------------- | ----------------------------------- | ------------------- |
| `--addr`            | チェック対象のHTTPエンドポイントURL | なし（必須）        |
| `--method`          | HTTPメソッド                        | `GET`               |
| `--timeout`         | リクエストのタイムアウト            | `1s`                |
| `--expected-status` | 正常と判断するHTTPステータスコード  | `200`               |
| `--tls`             | TLSを使用する                       | `false`             |
| `--insecure`        | TLS証明書の検証をスキップする       | `false`             |
| `--user-agent`      | User-Agentヘッダー                  | `http-health-probe` |

## 🐳 distrolessイメージでの利用例

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

## 📝 ライセンス

本プロジェクトはMITライセンスの下で公開されています。

## 🤝 貢献

IssueやPull Requestは歓迎します。バグ報告や機能提案もお気軽にどうぞ。
