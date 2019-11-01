package htmlparser_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/htmlparser"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorHtmlParserTypeConversion_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorHtmlParserTypeConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserTypeConversion_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorHtmlParserTypeConversion{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserParsing_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorHtmlParserParsing{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserParsing_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorHtmlParserParsing{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserNoSuchElement_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorHtmlParserNoSuchElement{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserNoSuchElement_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorHtmlParserNoSuchElement{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserTooManyElements_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorHtmlParserTooManyElements{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorHtmlParserTooManyElements_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorHtmlParserTooManyElements{Err: expectedStr}

	Compare(t, err, expectedStr)
}
