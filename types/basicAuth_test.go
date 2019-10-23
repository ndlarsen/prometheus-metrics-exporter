package types

import (
	"encoding/json"
	"fmt"
	. "prometheus-metrics-exporter/pmeerrors"
	"testing"
)

const (
	baUsername = "basicAuthUsername"
	baPassword = "basicAuthPassword"
)

func Test_BasicAuth_OK(t *testing.T) {
	var jsonString = fmt.Sprintf(`{"username": "%s", "password": "%s"}`, baUsername, baPassword)
	var jsonBasicAuthOk = []byte(jsonString)

	var ba BasicAuth

	err := json.Unmarshal(jsonBasicAuthOk, &ba)

	if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err.Error())
	} else if ba.Username != baUsername {
		t.Fatalf("Test failed unexpectedly: ba.Username mismatch")
	} else if ba.Password != baPassword {
		t.Fatalf("Test failed unexpectedly: ba.Password mismatch")
	}
}

func Test_BasicAuth_Invalid_JSON(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"username": "%s", "password" "%s}`, baUsername, baPassword)
	var jsonBytes = []byte(jsonString)

	var ba BasicAuth

	err := json.Unmarshal(jsonBytes, &ba)

	if err != nil && err == err.(*json.SyntaxError) {
		t.Logf("Test succeeded.")
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly")
	}
}

func Test_BasicAuth_Empty_Username(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"username": "", "password": "%s"}`, baPassword)
	var jsonLabelOK = []byte(jsonString)

	var ba BasicAuth

	err := json.Unmarshal(jsonLabelOK, &ba)

	if err != nil && err == err.(ErrorBasicAuthUnmarshal) {
		t.Logf("Test failed as expected: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err)
	} else {
		t.Fatalf("Test failed unexpectedly")
	}

}

func Test_BasicAuth_Empty_Password(t *testing.T) {

	var jsonString = fmt.Sprintf(`{"username": "%s", "password": ""}`, baPassword)
	var jsonLabelOK = []byte(jsonString)

	var ba BasicAuth

	err := json.Unmarshal(jsonLabelOK, &ba)

	if err != nil && err == err.(ErrorBasicAuthUnmarshal) {
		t.Logf("Test failed as expected: %s", err.Error())
	} else if err != nil {
		t.Fatalf("Test failed unexpectedly: %s", err)
	} else {
		t.Fatalf("Test failed unexpectedly")
	}

}
