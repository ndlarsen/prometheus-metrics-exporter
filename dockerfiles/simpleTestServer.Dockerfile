FROM golang:1.17.12-alpine3.16

WORKDIR /go/src/simpleTestServer
COPY . .

RUN GCO_ENABLED=0 GOOS=linux go build -o runTestServer ./test/e2e/run_test_server.go
CMD ["./runTestServer", "-port=8080"]