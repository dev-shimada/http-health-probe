FROM --platform=$BUILDPLATFORM golang:1.24.3-bookworm as vscode

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ENV LESSCHARSET utf-8

WORKDIR /workspace

RUN  <<EOF
apt-get update && apt-get install -y vim git unzip less bash-completion jq
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v github.com/cweill/gotests/gotests@latest
go install -v github.com/goreleaser/goreleaser/v2@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc
echo export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
EOF

COPY . /workspace
