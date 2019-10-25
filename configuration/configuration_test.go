package configuration_test

import (
	"encoding/json"
	. "prometheus-metrics-exporter/configuration"
	. "prometheus-metrics-exporter/pmeerrors"
	"prometheus-metrics-exporter/types"
	"reflect"
	"testing"
)

func TestInvalidFileName(t *testing.T) {
	var invalidFileName = "../test_related/invalidFileName.json"
	_, err := LoadConfig(invalidFileName)

	if recover() != nil {
		t.Errorf("Failed to panic on invalid filename.")
		return
	}

	if err == nil {
		t.Errorf("Didn't fail as expected: TestInvalidFileName.")
		return
	}

}

func TestInvalidFileFormat(t *testing.T) {
	var fileName = "../test_related/invalidJsonFormat.json"
	_, err := LoadConfig(fileName)

	if err == err.(ErrorConfigConversion) {
		t.Log("Could not convert file content as expected.")
	} else if err != nil && err != err.(ErrorConfigConversion) {
		t.Errorf("Failed unexpectedly.")
	} else {
		t.Errorf("Didn't fail as expected.")
	}
}

func TestValidFileFormat(t *testing.T) {

	const fullPath = "../test_related/validConfig.json"

	fileConfig, loadErr := LoadConfig(fullPath)

	if loadErr != nil {
		t.Fatalf("Unexpected failure. Failed to read file: %s", loadErr)
	}

	const jsonString = `{
  "scrapeTargets": [
    {
      "url": "https://jsonplaceholder.typicode.com/users",
      "metrics": [
        {
          "name": "name",
          "help": "help",
          "path": "json.path.01",
          "instrumentType": "gauge"
        },
        {
          "name": "name",
          "help": "help",
          "path": "json.path.02",
          "instrumentType": "counter"
        }
      ],
      "labels": [
        {"name": "LabelName", "value": "LabelValue"}
      ],
      "mimeType": "json",
      "jobName": "promName",
      "timeoutInSecs": 15
    },
    {
      "url": "https://jsonplaceholder.typicode.com/users",
      "metrics": [
        {
          "name": "name",
          "help": "help",
          "path": "json.path.01",
          "instrumentType": "gauge"
        },
        {
          "name": "name",
          "help": "help",
          "path": "json.path.02",
          "instrumentType": "counter"
        }
      ],
      "labels": [
        {"name": "LabelName", "value": "LabelValue"}
      ],
      "mimeType": "json",
      "jobName": "promName",
      "timeoutInSecs": 15
    }
  ],
  "pushGatewayUrl": "gateWayUrl"
}`

	var jsonBytes = []byte(jsonString)
	var stringConfig *types.Config

	err := json.Unmarshal(jsonBytes, &stringConfig)

	if err != nil {
		t.Fatalf("Unexpected failure. Unable to unmarshal json: %s", err)
	}

	if !reflect.DeepEqual(fileConfig, stringConfig) {
		t.Fatalf("Config structs are not equal")
	}

}
