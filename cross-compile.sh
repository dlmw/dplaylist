#!/bin/sh
for GOOS in darwin linux windows; do
  for GOARCH in amd64; do
    export GOOS GOARCH
    go build -v -o dplaylist-$GOOS-$GOARCH
  done
done