FROM golang:1.15.2-alpine3.12

# System setup
RUN apk update && apk upgrade \
    && apk add git curl build-base autoconf automake libtool \
    && apk add protoc protobuf-dev

# Install protoc-gen-go
ENV GO111MODULE=on
RUN go get github.com/golang/protobuf/protoc-gen-go \
    && go get github.com/gogo/protobuf/protoc-gen-gofast
ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /home/noname

# WIP (protoc-gen-go-grpc): https://github.com/grpc/grpc.io/issues/298
RUN git clone -b v1.31.0 https://github.com/grpc/grpc-go \
    && cd grpc-go/cmd/protoc-gen-go-grpc \
    && go install .

#No Root
RUN adduser -D -u 1000 noname -h  /home/noname
USER noname

RUN mkdir import
RUN mkdir export