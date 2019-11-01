package label_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/label"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorLabelUnmarshal_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorLabelUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorLabelUnmarshal_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorLabelUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}
