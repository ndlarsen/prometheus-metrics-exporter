FROM golang:1.11.13-alpine3.10

RUN apk add make git curl bash

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /go/src/build

COPY . .

RUN make build_linux
