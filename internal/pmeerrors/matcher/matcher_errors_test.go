package matcher_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/matcher"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorMatcherRegexCompileError_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorMatcherRegexCompileError{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMatcherRegexCompileError_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorMatcherRegexCompileError{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMatcherRegexNoMatch_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorMatcherRegexNoMatch{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMatcherRegexNoMatch_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorMatcherRegexNoMatch{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMatcherRegexNoCaptureGroup_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorMatcherRegexNoCaptureGroup{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorMatcherRegexNoCaptureGroup_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorMatcherRegexNoCaptureGroup{Err: expectedStr}

	Compare(t, err, expectedStr)
}
