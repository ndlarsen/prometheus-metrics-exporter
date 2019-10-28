FROM golang:1.11.13-alpine3.10

WORKDIR /go/src/simpleTestServer
COPY ./simpleTestServer .

RUN pwd
RUN ls -la
RUN GCO_ENABLED=0 GOOS=linux go build -o simpleTestServer ./main.go
CMD ["./simpleTestServer", "-port=8080"]
