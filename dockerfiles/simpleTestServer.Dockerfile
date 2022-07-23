FROM golang:1.15.15-alpine3.14

WORKDIR /go/src/simpleTestServer
COPY . .

RUN ls -la

RUN GCO_ENABLED=0 GOOS=linux go build -o runTestServer ./test/e2e/run_test_server.go
CMD ["./runTestServer", "-port=8080"]
