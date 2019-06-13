package types_test

import (
	"encoding/json"
	"fmt"
	. "prometheus-metrics-exporter/pmeerrors"
	. "prometheus-metrics-exporter/types"
	"testing"
)

const (
	lName  = "labelName"
	lValue = "labelValue"
)

func Test_Label_OK(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)
	var jsonLabelOK = []byte(jsonString)

	var lbl Label

	err := json.Unmarshal(jsonLabelOK, &lbl)

	t.Logf("%+v", lbl)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if lbl.Name != lName {
		t.Fatalf("Test failed unexpectedly: lbl.Name mismatch")
	} else if lbl.Value != lValue {
		t.Fatalf("Test failed unexpectedly: lbl.Value mismatch")
	}

}

func Test_Label_Invalid_JSON(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"name": "%s", "value" "%s}`, lName, lValue)
	var jsonBytes = []byte(jsonString)

	var lbl Label

	//err := json.Unmarshal(jsonBytes, &lbl)
	err := lbl.UnmarshalJSON(jsonBytes)

	if err == nil {
		t.Fatalf("Test failed unexpectedly")
	} else if err != nil {
		t.Logf("Test failed as expected: %s", err.Error())
	}

}

func Test_Label_Empty_Name(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"name": "", "value": "%s"}`, lValue)
	var jsonLabelOK = []byte(jsonString)

	var lbl Label

	err := json.Unmarshal(jsonLabelOK, &lbl)

	if err != nil && err == err.(ErrorLabelUnmarshal) {
		t.Logf("Test failed as expected: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err)
	} else {
		t.Fatalf("Test failed unexpectedly")
	}

}

func Test_Label_Empty_Value(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"name": "%s", "value": ""}`, lName)
	var jsonLabelOK = []byte(jsonString)

	var lbl Label

	err := json.Unmarshal(jsonLabelOK, &lbl)

	if err != nil && err == err.(ErrorLabelUnmarshal) {
		t.Logf("Test failed as expected: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err)
	} else {
		t.Fatalf("Test failed unexpectedly")
	}

}
