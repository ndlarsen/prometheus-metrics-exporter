#!/bin/bash

exit_error() {
    if [ $# -ne 0 ]
    then
      echo "$@"
    fi
    exit 1
}

echo "Running PME..."
./binaries/prometheus-metrics-exporter-linux-amd64 -config="./test/e2e/test_config.json" || exit_error "PME failed to run"

echo "Running e2e tests..."

go test -v -failfast ./test/e2e/e2e_test.go
