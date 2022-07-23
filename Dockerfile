FROM golang:1.12.17-alpine3.11

RUN apk add make git

ENV GO111MODULE=on \
    CGO_ENABLED=0

ARG UID=1000
ARG GID=1000

RUN addgroup --gid $GID --system builduser && adduser -u $UID -D -G builduser builduser
USER builduser
WORKDIR /home/builduser
