package jsonparser_test

import (
	"fmt"
	. "prometheus-metrics-exporter/internal/jsonparser"
	"prometheus-metrics-exporter/internal/pmeerrors/jsonparser"
	"prometheus-metrics-exporter/internal/pmeerrors/matcher"
	"testing"
)

const jsonString = `{
  "id": 1,
  "boolValue": true,
  "notPresent": null,
  "mixedAttribute": "It seems that 42 is a special number. ",
  "address": {
    "street": "Kulas Light",
    "geo": {
      "lat": "-37.3159",
    }
  },
}`

func Test_FetchValue_OK_String_With_Number(t *testing.T) {

	const testStr = "address.geo.lat"
	const expectedValue = -37.3159

	value, err := FetchValue(testStr, []byte(jsonString), "")

	if err == nil && value == expectedValue {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%f\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}
}

func Test_FetchValue_OK_Number(t *testing.T) {

	const testStr = "id"
	const expectedValue = 1

	value, err := FetchValue(testStr, []byte(jsonString), "")

	if err == nil && value == expectedValue {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%d\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValue_NonExistingValue(t *testing.T) {

	const testStr = "bogus.path"

	_, err := FetchValue(testStr, []byte(jsonString), "")

	if err != nil {
		t.Log("Failed as expected.")
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValue_StringChar(t *testing.T) {

	const testStr = "address.street"

	_, err := FetchValue(testStr, []byte(jsonString), "")

	if err != nil {
		t.Log("Failed as expected.")
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValue_Bool(t *testing.T) {

	const testStr = "boolValue"
	const expectedValue = true

	_, err := FetchValue(testStr, []byte(jsonString), "")

	if err != nil && err == err.(jsonparser.ErrorJsonParserInvalidType) {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%t\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValue_Nil(t *testing.T) {

	const testStr = "notPresent"

	_, err := FetchValue(testStr, []byte(jsonString), "")

	if err != nil && err == err.(jsonparser.ErrorJsonParserInvalidType) {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"nil\" as expected.", testStr)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValue_Regex_OK(t *testing.T) {

	const path = "mixedAttribute"
	const pattern = `that (\d+) is`

	var expectedResult float64 = 42

	result, parseErr := FetchValue(path, []byte(jsonString), pattern)

	if parseErr == nil && result == expectedResult {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_Compile_Error(t *testing.T) {

	const path = "mixedAttribute"
	const pattern = `(that \d+ is`

	result, parseErr := FetchValue(path, []byte(jsonString), pattern)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexCompileError) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", result, "Error: ", parseErr)
	} else {
		t.Log("value: ", result, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Match(t *testing.T) {

	const path = "address.street"
	const pattern = `that (\d+) is`

	result, parseErr := FetchValue(path, []byte(jsonString), pattern)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexNoMatch) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", result, "Error: ", parseErr)
	} else {
		t.Log("value: ", result, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Capture_Group(t *testing.T) {

	const path = "mixedAttribute"
	const pattern = `that \d+ is`

	result, parseErr := FetchValue(path, []byte(jsonString), pattern)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexNoCaptureGroup) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", result, "Error: ", parseErr)
	} else {
		t.Log("value: ", result, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}
