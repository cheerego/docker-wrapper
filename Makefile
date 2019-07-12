SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.io

# build with verison infos
versionDir = "github.com/cheerego/docker-wrapper/cmd"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.GitTag=${gitTag} -X ${versionDir}.BuildDate=${buildDate} -X ${versionDir}.GitCommit=${gitCommit} -X ${versionDir}.GitTreeState=${gitTreeState}"



all:clean fmt
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/amd64/darwin/docker-wrapper -v -ldflags ${ldflags} .
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/amd64/linux/docker-wrapper -v -ldflags ${ldflags} .
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/amd64/windows/docker-wrapper.exe -v -ldflags ${ldflags} .
clean:
	rm -rf dist
fmt:
	gofmt -w .

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"