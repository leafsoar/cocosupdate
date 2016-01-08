#!/bin/bash

set -eu

echo "go build ..."
cd ../
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/cocosupdate
cd -
docker build --force-rm=true -t leafsoar/cocosupdate .
rm cocosupdate

