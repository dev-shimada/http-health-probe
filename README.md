# http-health-probe

[![CI](https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml/badge.svg)](https://github.com/dev-shimada/http-health-probe/actions/workflows/CI.yaml)
[![Coverage Status](https://coveralls.io/repos/github/your-org/http-health-probe/badge.svg?branch=main)](https://coveralls.io/github/your-org/http-health-probe?branch=main)

http-health-probeは、HTTPベースのAPIのヘルスチェックを行うためのシンプルなコマンドラインツールです。

## ✨ 特徴

- C言語のライブラリ不要（Goのみで実装）
- distrolessイメージで動作可能
- シンプルなコマンドラインインターフェース

## 📦 インストール

### バイナリダウンロード

リリースページから各プラットフォーム向けのバイナリをダウンロードしてください。

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

| オプション      | 説明                                         | デフォルト値         |
|----------------|----------------------------------------------|---------------------|
| `--addr`        | チェック対象のHTTPエンドポイントURL           | なし（必須）        |

## 🐳 distrolessイメージでの利用例

```Dockerfile
FROM gcr.io/distroless/base-debian12:latest
COPY --chown=nonroot:nonroot --from=build http-health-probe /bin/http-health-probe
HEALTHCHECK --interval=10s --timeout=3s --start-period=5s CMD ["/bin/http-health-probe", "--addr=http://localhost:3000"]
```

## 📝 ライセンス

本プロジェクトはMITライセンスの下で公開されています。

## 🤝 貢献

IssueやPull Requestは歓迎します。バグ報告や機能提案もお気軽にどうぞ。
