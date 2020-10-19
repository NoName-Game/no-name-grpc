FROM golang:1.15.2-alpine3.12

# System setup
RUN apk update && apk upgrade \
    && apk add git curl build-base autoconf automake libtool \
    && apk add protoc protobuf-dev

# Install protoc-gen-go (GOGO)
RUN go get github.com/gogo/protobuf/protoc-gen-gogo \
    && go get github.com/gogo/protobuf/proto \
    && go get github.com/gogo/protobuf/gogoproto \
    && go get github.com/gogo/protobuf/protoc-gen-gofast \
    && go get github.com/gogo/protobuf/protoc-gen-gogofast

ENV PATH="$PATH:$(go env GOPATH)/bin"

RUN echo $GOPATH

WORKDIR /home/noname

#No Root
RUN adduser -D -u 1000 noname -h  /home/noname
USER noname

RUN mkdir import
RUN mkdir export