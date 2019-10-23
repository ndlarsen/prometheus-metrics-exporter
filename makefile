# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o prometheus-metrics-exporter main.go

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_ARCH=amd64
PLATFORMS := linux/$(GO_ARCH) windows/$(GO_ARCH)
BINARY_NAME=prometheus-metrics-exporter

all: test release
build_linux:
	 GOOS=linux GO_ARCH=${GO_ARCH} $(GO_BUILD) -o $(BINARY_NAME)-linux-${GO_ARCH} -v
build_win:
	GOOS=linux GO_ARCH=${GO_ARCH} $(GO_BUILD) -o $(BINARY_NAME)-windows-${GO_ARCH} -v
release:
	build_linux build_win
test:
	$(GO_TEST) -v -failfast ./...
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME)-windows
deps:
	$(GO_GET) github.com/prometheus/client_golang v0.9.1
	$(GO_GET) github.com/tidwall/gjson v1.1.3
	$(GO_GET) github.com/antchfx/htmlquery v0.0.0-20180925020018-98389addba3d
