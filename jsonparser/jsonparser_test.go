package jsonparser_test

import (
	"fmt"
	"prometheus-metrics-exporter/jsonparser"
	. "prometheus-metrics-exporter/pmeerrors"
	"testing"
)

const jsonString = `{
  "id": 1,
  "boolValue": true,
  "notPresent": null,
  "address": {
    "street": "Kulas Light",
    "geo": {
      "lat": "-37.3159",
    }
  },
}`

func Test_FetchValuesFromJSON_OK_String_With_Number(t *testing.T) {

	const testStr = "address.geo.lat"
	const expectedValue = -37.3159

	value, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err == nil && value == expectedValue {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%f\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}
}

func Test_FetchValuesFromJSON_OK_Number(t *testing.T) {

	const testStr = "id"
	const expectedValue = 1

	value, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err == nil && value == expectedValue {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%d\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValuesFromJSON_NonExistingValue(t *testing.T) {

	const testStr = "bogus.path"

	_, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err != nil {
		t.Log("Failed as expected.")
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValuesFromJSON_StringChar(t *testing.T) {

	const testStr = "address.street"

	_, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err != nil {
		t.Log("Failed as expected.")
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValuesFromJSON_Bool(t *testing.T) {

	const testStr = "boolValue"
	const expectedValue = true

	_, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err != nil && err == err.(ErrorJsonParserInvalidType) {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"%t\" as expected.", testStr, expectedValue)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}

func Test_FetchValuesFromJSON_Nil(t *testing.T) {

	const testStr = "notPresent"

	_, err := jsonparser.FetchValue(testStr, []byte(jsonString))

	if err != nil && err == err.(ErrorJsonParserInvalidType) {
		str := fmt.Sprintf("Extracted value of \"%s\" was \"nil\" as expected.", testStr)
		t.Log(str)
	} else {
		t.Fatalf("Test did not succeed as expected.")
	}

}
