package configuration_test

import (
	"encoding/json"
	. "prometheus-metrics-exporter/internal/configuration"
	"prometheus-metrics-exporter/internal/pmeerrors/config"
	. "prometheus-metrics-exporter/internal/types"
	"reflect"
	"testing"
)

func TestInvalidFileName(t *testing.T) {
	var invalidFileName = "../../test/invalidFileName.json"
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
	var fileName = "../../test/unit/invalidJsonFormat.json"
	_, err := LoadConfig(fileName)

	if err == err.(config.ErrorConfigConversion) {
		t.Log("Could not convert file content as expected.")
	} else if err != nil && err != err.(config.ErrorConfigConversion) {
		t.Errorf("Failed unexpectedly.")
	} else {
		t.Errorf("Didn't fail as expected.")
	}
}

func TestValidFileFormat(t *testing.T) {

	const fullPath = "../../test/unit/validConfig.json"

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
	var stringConfig *Config

	err := json.Unmarshal(jsonBytes, &stringConfig)

	if err != nil {
		t.Fatalf("Unexpected failure. Unable to unmarshal json: %s", err)
	}

	if !reflect.DeepEqual(fileConfig, stringConfig) {
		t.Fatalf("Config structs are not equal")
	}

}
