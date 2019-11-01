package config_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/config"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorConfigReadFile_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorConfigConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorConfigReadFile_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorConfigConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorConfigConversion_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorConfigConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorConfigConversion_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorConfigConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorConfigUnmarshal_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorConfigUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorConfigUnmarshal_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorConfigUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}
