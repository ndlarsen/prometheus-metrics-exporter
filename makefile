# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
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

clean:
	bash -c "rm -rfv binaries/*"

all_local: test_unit_local build_binary

build_binary:
	GOOS=linux GO_ARCH=${GO_ARCH} ${GO_BUILD} -o ${BUILD_TARGET_DIR}/$(BINARY_NAME)-linux-${GO_ARCH} -v cmd/main.go

build_binary_docker:
	${DOCKER_BUILD}
	${DOCKER_RUN} make build_binary

test_unit_local:
	$(GO_TEST) -v -failfast ./internal/...

test_unit_docker:
	${DOCKER_BUILD}
	${DOCKER_RUN} make test_unit_local

test_e2e:
	docker-compose -f ./e2etest/e2etest.docker-compose.yml up --build --abort-on-container-exit

all: build_binary clean build_binary_docker clean test_unit_local test_unit_docker test_e2e

