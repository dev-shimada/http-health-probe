#!/bin/bash
set -e

# tools
echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc
echo export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
