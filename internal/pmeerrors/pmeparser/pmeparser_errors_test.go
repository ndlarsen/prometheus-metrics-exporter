package pmeparser_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/pmeparser"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorParserInvalidContentType_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorParserInvalidContentType{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorParserInvalidContentType_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorParserInvalidContentType{Err: expectedStr}

	Compare(t, err, expectedStr)
}
