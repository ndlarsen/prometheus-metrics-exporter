FROM golang:1.16.15-alpine3.15

RUN apk add make git curl bash

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /go/src/build

COPY . .

RUN make build_binary
