package request_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/request"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorRequestClient_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestClient{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestClient_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestClient{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestTimeOut_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestTimeOut{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestTimeOut_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestTimeOut{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus401_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestResponseStatus401{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus401_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestResponseStatus401{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus403_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestResponseStatus403{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus403_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestResponseStatus403{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus404_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestResponseStatus404{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus404_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestResponseStatus404{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus500_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestResponseStatus500{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatus500_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestResponseStatus500{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatusNot200_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestResponseStatusNot200{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestResponseStatusNot200_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestResponseStatusNot200{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestInvalidContentTypeFound_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestInvalidContentTypeFound{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestInvalidContentTypeFound_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestInvalidContentTypeFound{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestUnableToReadBody_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestUnableToReadBody{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestUnableToReadBody_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestUnableToReadBody{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestContentTypeParse_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorRequestContentTypeParse{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorRequestContentTypeParse_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorRequestContentTypeParse{Err: expectedStr}

	Compare(t, err, expectedStr)
}
