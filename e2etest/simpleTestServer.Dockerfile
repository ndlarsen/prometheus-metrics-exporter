FROM golang:1.14.15-alpine3.13

WORKDIR /go/src/simpleTestServer
COPY ./simpleTestServer .

RUN pwd
RUN ls -la
RUN GCO_ENABLED=0 GOOS=linux go build -o simpleTestServer ./main.go
CMD ["./simpleTestServer", "-port=8080"]
