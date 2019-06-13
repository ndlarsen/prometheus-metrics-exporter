package types_test

import (
	"encoding/json"
	"fmt"
	. "prometheus-metrics-exporter/pmeerrors"
	. "prometheus-metrics-exporter/types"
	"reflect"
	"testing"
)

const confPushGatewayUrl = "confTestUrl"

func Test_Config_OK(t *testing.T) {

	var st = ScrapeTarget{
		Url: stUrl,
		Metrics: []Metric{
			{
				Name:           mName,
				Help:           mHelp,
				InstrumentType: mInstrumentType,
				Path:           mPath,
				Regex:          mRegex,
			},
		},
		Labels: []Label{
			{
				Name:  lName,
				Value: lValue,
			},
		},
		MimeType:      stMimeType,
		JobName:       stJobName,
		TimeoutInSecs: stTimeOutInSecs,
	}

	var confStruct = Config{
		ScrapeTargets: []ScrapeTarget{
			st,
		},
		PushGatewayUrl: confPushGatewayUrl,
	}

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
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	confJsonString := fmt.Sprintf(`{
	"scrapeTargets": [%s],
	"pushGatewayUrl":"%s"
}`, stJsonString, confPushGatewayUrl)

	var jsonBytes = []byte(confJsonString)

	var conf Config

	err := json.Unmarshal(jsonBytes, &conf)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if !reflect.DeepEqual(conf, confStruct) {
		t.Fatalf("Test failed unexpectedly. Structs are not equal")
	} else if len(conf.ScrapeTargets) != len(confStruct.ScrapeTargets) && len(conf.ScrapeTargets) != 1 {
		t.Fatalf("Test failed unexpectedly: Url mismatch")
	} else if len(conf.ScrapeTargets[0].Metrics) != 1 || len(confStruct.ScrapeTargets[0].Metrics) != 1 ||
		len(conf.ScrapeTargets[0].Metrics) != len(confStruct.ScrapeTargets[0].Metrics) ||
		!reflect.DeepEqual(conf.ScrapeTargets[0].Metrics, confStruct.ScrapeTargets[0].Metrics) {
		t.Fatalf("Test failed unexpectedly: Metrics mismatch")
	} else if len(conf.ScrapeTargets[0].Labels) != len(confStruct.ScrapeTargets[0].Labels) ||
		!reflect.DeepEqual(conf.ScrapeTargets[0].Labels, confStruct.ScrapeTargets[0].Labels) {
		t.Fatalf("Test failed unexpectedly: Labels mismatch")
	} else if conf.PushGatewayUrl != confStruct.PushGatewayUrl {
		t.Fatalf("Test failed unexpectedly: MimeType mismatch")
	} else {
		t.Logf("Test succeeded.")
	}

}

func Test_Config_Invalid_JSON(t *testing.T) {

	confJsonString := fmt.Sprintf(`{
	"scrapeTargets": []
	"pushGatewayUrl":"%s
}`, confPushGatewayUrl)

	var jsonBytes = []byte(confJsonString)

	var conf Config

	//err := json.Unmarshal(jsonBytes, &conf)
	err := conf.UnmarshalJSON(jsonBytes)

	if err != nil && err == err.(*json.SyntaxError) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else {
		t.Fatal("Test failed unexpectedly.")
	}

}

func Test_Config_Empty_PushGatewayUrl(t *testing.T) {

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
	"jobName": "%s",
	"timeoutInSecs": %d
}`, stUrl, metricJsonString, lblJsonString, stMimeType, stJobName, stTimeOutInSecs)

	confJsonString := fmt.Sprintf(`{
	"scrapeTargets": [%s],
	"pushGatewayUrl":""
}`, stJsonString)

	var jsonBytes = []byte(confJsonString)

	var conf Config

	err := json.Unmarshal(jsonBytes, &conf)

	if err != nil && err == err.(ErrorConfigUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else {
		t.Fatal("Test failed unexpectedly.")
	}

}

func Test_Config_Empty_ScrapeTargets(t *testing.T) {

	confJsonString := fmt.Sprintf(`{
	"scrapeTargets": [],
	"pushGatewayUrl":"%s"
}`, confPushGatewayUrl)

	var jsonBytes = []byte(confJsonString)

	var conf Config

	err := json.Unmarshal(jsonBytes, &conf)

	if err != nil && err == err.(ErrorConfigUnmarshal) {
		t.Logf("Test succeeded: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else {
		t.Fatal("Test failed unexpectedly.")
	}

}
