#!/bin/bash
export APP_VERSION=$(git describe --tags --abbrev=0)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s -X main.version=$APP_VERSION" -o bin/$1 -v main.go