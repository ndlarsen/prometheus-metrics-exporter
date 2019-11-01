package basicauth_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/basicauth"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorBasicAuthUnmarshal_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorBasicAuthUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorBasicAuthUnmarshal_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorBasicAuthUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)

}
