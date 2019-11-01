package types_test

import (
	"encoding/json"
	"fmt"
	"prometheus-metrics-exporter/internal/pmeerrors/metric"
	. "prometheus-metrics-exporter/internal/types"
	"testing"
)

const (
	mName           = "metricName"
	mHelp           = "metricHelp"
	mPath           = "metricPath"
	mInstrumentType = "metricInstrumentType"
	mRegex          = "metricRegex"
)

func Test_Metric_OK(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if m.Name != mName {
		t.Fatalf("Test failed unexpectedly: Name mismatch")
	} else if m.Help != mHelp {
		t.Fatalf("Test failed unexpectedly: Help mismatch")
	} else if m.Path != mPath {
		t.Fatalf("Test failed unexpectedly: Path mismatch")
	} else if m.InstrumentType != mInstrumentType {
		t.Fatalf("Test failed unexpectedly: InstrumentType mismatch")
	} else if m.Regex != mRegex {
		t.Logf("Regex: %s: %s", mRegex, m.Regex)
		t.Fatalf("Test failed unexpectedly: Regex mismatch")
	} else {
		t.Logf("Test succeeded.")
	}
}

func Test_Metric_Empty_Name(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mHelp, mPath, mInstrumentType, mRegex)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil && err == err.(metric.ErrorMetricUnmarshal) {
		t.Logf("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly")
	}
}

func Test_Metric_Empty_Help(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mPath, mInstrumentType, mRegex)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil && err == err.(metric.ErrorMetricUnmarshal) {
		t.Logf("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly")
	}
}

func Test_Metric_Empty_Path(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mInstrumentType, mRegex)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil && err == err.(metric.ErrorMetricUnmarshal) {
		t.Logf("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly")
	}
}

func Test_Metric_Empty_InstrumentType(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mRegex)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil && err == err.(metric.ErrorMetricUnmarshal) {
		t.Logf("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly")
	}
}

func Test_Metric_Omitted_Regex(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"InstrumentType": "%s",
	"Path": "%s"
}`, mName, mHelp, mInstrumentType, mPath)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else {
		t.Logf("Test succeeded")
	}
}

func Test_Metric_Empty_String_Regex(t *testing.T) {
	var jsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"InstrumentType": "%s",
	"Regex": "",
	"Path": "%s"
}`, mName, mHelp, mInstrumentType, mPath)

	var jsonBytes = []byte(jsonString)

	var m Metric

	err := json.Unmarshal(jsonBytes, &m)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if m.Regex != "" {
		t.Fatalf("Test failed unexpectedly: Regex is not an empty string")
	} else {
		t.Logf("Test succeeded")
	}
}

func Test_Metric_Json_Invalid_JSON(t *testing.T) {
	var invalidJsonString = fmt.Sprintf(`{
	"Name: "%s",
	"Help": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mInstrumentType, mRegex)

	var jsonBytes = []byte(invalidJsonString)

	var m Metric

	err := m.UnmarshalJSON(jsonBytes)

	if err != nil && err == err.(*json.SyntaxError) {
		t.Logf("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly")
	}
}
