#!/bin/sh

path="target"
rm -rf $path
mkdir $path
mkdir $path/config
cp -f ./config/config.toml $path/config/config.toml
cp -f ./config/gdao.toml $path/config/gdao.toml
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./launcher.go
mv ./launcher $path/launcher
tar zcvf $path.tar.gz $path