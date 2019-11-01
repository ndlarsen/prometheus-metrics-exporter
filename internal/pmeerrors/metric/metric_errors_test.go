package metric_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/metric"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorMetricUnmarshal_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorMetricUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMetricUnmarshal_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorMetricUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}
