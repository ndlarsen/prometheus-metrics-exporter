package types_test

import (
	"encoding/json"
	"fmt"
	. "prometheus-metrics-exporter/pmeerrors"
	. "prometheus-metrics-exporter/types"
	"reflect"
	"testing"
)

var (
	stUrl     = "testStUrl"
	stMetrics = []Metric{
		{
			Name:           mName,
			Help:           mHelp,
			InstrumentType: mInstrumentType,
			Path:           mPath,
			Regex:          mRegex,
		},
	}
	stLabels = []Label{
		{
			Name:  lName,
			Value: lValue,
		},
	}
	stMimeType      = "json"
	stJobName       = "scJobName"
	stTimeOutInSecs = 10
)

func Test_ScrapeTarget_OK(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := st.UnmarshalJSON(jsonBytes)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if st.Url != stUrl {
		t.Fatalf("Test failed unexpectedly: Url mismatch")
	} else if len(st.Metrics) != len(stMetrics) || !reflect.DeepEqual(st.Metrics, stMetrics) {
		t.Fatalf("Test failed unexpectedly: Metrics mismatch")
	} else if len(st.Labels) != len(stLabels) || !reflect.DeepEqual(st.Labels, stLabels) {
		t.Fatalf("Test failed unexpectedly: Labels mismatch")
	} else if st.MimeType != stMimeType {
		t.Fatalf("Test failed unexpectedly: MimeType mismatch")
	} else if st.JobName != stJobName {
		t.Fatalf("Test failed unexpectedly: JobName mismatch")
	} else if st.TimeoutInSecs != stTimeOutInSecs {
		t.Fatalf("Test failed unexpectedly: TimeOutInSecs mismatch")
	} else {
		t.Logf("Test succeeded.")
	}

}

func Test_ScrapeTarget_Invalid_JSON(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "%s"
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	//err := json.Unmarshal(jsonBytes, &st)
	err := st.UnmarshalJSON(jsonBytes)

	if err != nil && err == err.(*json.SyntaxError) {
		t.Logf("Test succeeded.")
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_Url(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "%s",
	"timeoutInSecs": %d
}`, metricJsonString, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_Metrics(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_Labels(t *testing.T) {

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [],
	"mimeType": "%s",
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, stMimeType, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_MimeType(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "",
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stJobName, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_JobName(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"Name": "%s",
	"Help": "%s",
	"Path": "%s",
	"InstrumentType": "%s",
	"Regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stTimeOutInSecs)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}

func Test_ScrapeTarget_Empty_Timeout(t *testing.T) {

	var lblJsonString = fmt.Sprintf(`{"name": "%s", "value": "%s"}`, lName, lValue)

	var metricJsonString = fmt.Sprintf(`{
	"name": "%s",
	"help": "%s",
	"path": "%s",
	"instrumentType": "%s",
	"regex": "%s"
}`, mName, mHelp, mPath, mInstrumentType, mRegex)

	var stJsonString = fmt.Sprintf(`{
	"url": "%s",
	"metrics": [%s],
	"labels": [%s],
	"mimeType": "%s",
	"jobName": "%s"
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stJobName)

	var jsonBytes = []byte(stJsonString)

	var st ScrapeTarget

	err := json.Unmarshal(jsonBytes, &st)

	if err != nil && err == err.(ErrorScrapeTargetUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	}

}
