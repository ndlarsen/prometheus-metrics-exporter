FROM golang:1.11.13-alpine3.10

RUN apk add make git curl bash

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build_dir

COPY . .

#RUN make build_linux

#RUN chmod +x ./test_related/run_pme_e2e_tests.sh
#RUN ./test_related/run_pme_e2e_tests.sh

#RUN go test -v --failfast ./e2etest/e2e_test.go
