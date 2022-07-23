FROM golang:1.17.12-alpine3.16

RUN apk add make git curl bash

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /go/src/build

COPY . .

RUN make build_binary
