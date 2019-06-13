# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o prometheus-metrics-exporter main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOARCH=amd64
PLATFORMS := linux/$(GOARCH) windows/$(GOARCH)
BINARY_NAME=prometheus-metrics-exporter

all: test realease
build_linux:
	 GOOS=linux GOARCH=${GOARCH} $(GOBUILD) -o $(BINARY_NAME)-linux -v
build_win:
	GOOS=linux GOARCH=${GOARCH} $(GOBUILD) -o $(BINARY_NAME)-windows -v
release:
	build_linux build_win
test:
	$(GOTEST) -v -failfast ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME)-windows
deps:
	$(GOGET) github.com/prometheus/client_golang v0.9.1
	$(GOGET) github.com/tidwall/gjson v1.1.3
	$(GOGET) github.com/antchfx/htmlquery v0.0.0-20180925020018-98389addba3d
