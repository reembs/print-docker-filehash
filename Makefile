all: mac linux

deps:
	go get \
	    github.com/docker/docker/pkg/tarsum \
	    github.com/docker/docker/pkg/archive \
	    github.com/docker/docker/pkg/pools

mac: deps
	GOARCH=amd64 GOOS=darwin go build -o bin/macos/print-docker-filehash main.go

linux: deps
	GOARCH=amd64 GOOS=linux go build -o bin/linux/print-docker-filehash main.go
