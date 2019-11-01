package instrument_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/instrument"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorInstrumentMissingValue_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorInstrumentMissingValue{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentMissingValue_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorInstrumentMissingValue{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentUnsupportedType_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorInstrumentUnsupportedType{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentUnsupportedType_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorInstrumentUnsupportedType{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentUrlParse_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorInstrumentUrlParse{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentUrlParse_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorInstrumentUrlParse{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentPushFailed_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorInstrumentPushFailed{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorInstrumentPushFailed_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorInstrumentPushFailed{Err: expectedStr}

	Compare(t, err, expectedStr)
}
