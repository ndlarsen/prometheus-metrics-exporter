FROM golang:1.16.15-alpine3.15

WORKDIR /go/src/simpleTestServer
COPY . .

RUN ls -la

RUN GCO_ENABLED=0 GOOS=linux go build -o runTestServer ./test/e2e/run_test_server.go
CMD ["./runTestServer", "-port=8080"]
