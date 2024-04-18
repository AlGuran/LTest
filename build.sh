#!/bin/bash
ROOTCMD=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
ROOT="${ROOTCMD}"
cd $ROOT/src
GOOS=linux  GOARCH=amd64 CGO_ENABLED=0 go build -o "$ROOT/bin/go_build_ltest" main.go
