package matcher_test

import (
	. "prometheus-metrics-exporter/internal/matcher"
	"prometheus-metrics-exporter/internal/pmeerrors/matcher"
	"testing"
)

// ([0-9]+)Mb

func Test_Matcher_Regex_OK(t *testing.T) {

	const str = `Vestibulum 650Mb auctor dapibus neque.`

	regex := `([0-9]+)Mb`
	var expectedResult string = "650"

	result, parseErr := Match(str, regex)

	if parseErr == nil && result == expectedResult {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_Compile_Error(t *testing.T) {

	const str = `67 Vestibulum 650Mb auctor dapibus 12 neque.`

	regex := `(\d+Mb`

	value, parseErr := Match(str, regex)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexCompileError) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Match(t *testing.T) {

	const str = `Vestibulum auctor dapibus neque.`

	regex := `([0-9]+)`

	value, parseErr := Match(str, regex)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexNoMatch) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}

func Test_FetchValue_Regex_No_Capture_Group(t *testing.T) {

	const str = `Vestibulum auctor 600Mb dapibus neque.`

	regex := `[0-9]+`

	value, parseErr := Match(str, regex)

	if parseErr != nil && parseErr == parseErr.(matcher.ErrorMatcherRegexNoCaptureGroup) {
		t.Log("Test succeeded as expected.")
		t.Log("value: ", value, "Error: ", parseErr)
	} else {
		t.Log("value: ", value, "Error: ", parseErr)
		t.Fatalf("Test failed unexpectedly: %s", parseErr)
	}

}
