#!/usr/bin/env bash
set -e

cd gopath/src/github.com/cloudfoundry/yagnats

go build -v ./...

./bin/test
