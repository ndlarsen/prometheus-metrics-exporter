# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o prometheus-metrics-exporter main.go

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) mod download
GO_ARCH=amd64
BUILD_TARGET_DIR=binaries
PLATFORMS := linux/$(GO_ARCH) windows/$(GO_ARCH)
BINARY_NAME=prometheus-metrics-exporter
UID=$$(id -u)
GID=$$(id -g)
DOCKER_TAG=pme_build:latest
DOCKER_BUILD=docker build --compress --build-arg UID=${UID} --build-arg GID=${GID} --build-arg GOARCH=${GO_ARCH} -t ${DOCKER_TAG} .
DOCKER_RUN=docker run --rm -it -v $$(pwd):/home/builduser ${DOCKER_TAG}

all_local: test_local build_local_release
clean:
	$(GO_CLEAN)
	rm -rf ${BUILD_TARGET_DIR}
build_linux:
	GOOS=linux GO_ARCH=${GO_ARCH} ${GO_BUILD} -o ${BUILD_TARGET_DIR}/$(BINARY_NAME)-linux-${GO_ARCH} -v
build_win:
	GOOS=linux GO_ARCH=${GO_ARCH} ${GO_BUILD} -o ${BUILD_TARGET_DIR}/$(BINARY_NAME)-windows-${GO_ARCH} -v
build_release_local: clean build_linux build_win
test_local:
	$(GO_TEST) -v -failfast ./...
test_docker:
	${GO_CLEAN}
	${DOCKER_BUILD}
	${DOCKER_RUN} make test_local
build_release_docker:
	${GO_CLEAN}
	${DOCKER_BUILD}
	${DOCKER_RUN} make build_release_local
deps:
	$(GO_GET) github.com/prometheus/client_golang v0.9.1
	$(GO_GET) github.com/tidwall/gjson v1.1.3
	$(GO_GET) github.com/antchfx/htmlquery v0.0.0-20180925020018-98389addba3d
