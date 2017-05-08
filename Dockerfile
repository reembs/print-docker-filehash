FROM golang:latest

VOLUME /go/result

RUN go get github.com/docker/docker/pkg/tarsum
RUN go get github.com/docker/docker/pkg/archive

COPY main.go /go

RUN go build main.go

CMD mv ./main /go/result/print-docker-filehash