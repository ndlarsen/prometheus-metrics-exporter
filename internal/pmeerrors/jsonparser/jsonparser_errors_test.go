package jsonparser_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/jsonparser"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorJsonParserInvalidType_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorJsonParserInvalidType{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorJsonParserInvalidType_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorJsonParserInvalidType{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorJsonParserTypeConversion_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorJsonParserTypeConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorJsonParserTypeConversion_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorJsonParserTypeConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorJsonParserValueEmpty_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorJsonParserValueEmpty{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorJsonParserValueEmpty_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorJsonParserValueEmpty{Err: expectedStr}

	Compare(t, err, expectedStr)
}
