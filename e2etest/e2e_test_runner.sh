#!/bin/bash

exit_error() {
    if [ $# -ne 0 ]
    then
      echo "$@"
    fi
    exit 1
}

echo "Running PME..."
./binaries/prometheus-metrics-exporter-linux-amd64 -config="./e2etest/e2e_test_config.json" || exit_error "PME failed to run"

echo "Running e2e tests..."

go test -v -failfast ./e2etest/e2e_test.go
