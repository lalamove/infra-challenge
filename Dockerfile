FROM golang:1.9.2
RUN apt-get update
RUN apt-get install -y vim curl ca-certificates

COPY . /go/src/code/
WORKDIR /go/src/code/

RUN export GOPATH=/go

RUN export GOBIN=/go/bin

CMD ["go","run","main.go"]
