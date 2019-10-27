package pmeparser_test

import (
	. "prometheus-metrics-exporter/internal/pmeerrors"
	. "prometheus-metrics-exporter/internal/pmeparser"
	"testing"
)

func Test_FetchValue_Json(t *testing.T) {

	url := ""
	path := "some.path"
	data := []byte(`{"some": {"path": 42}}`)
	contentType := "json"
	regex := ""
	var expectedValue float64 = 42

	actualValue, err := FetchValue(url, path, data, contentType, regex)

	if err == nil && actualValue == expectedValue {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly. %s", err.Error())
	}

}

func Test_FetchValue_Html(t *testing.T) {

	data := []byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>lorem ipsum</title>
</head>
<body>
<ul>
    <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
    <li>Aliquam tincidunt mauris eu risus.</li>
    <li>Vestibulum auctor 600Mb dapibus neque.</li>
</ul>
</body>
</html>`)

	url := ""
	path := "html/body/ul/li[3]"
	contentType := "html"
	regex := `(\d+)Mb`
	var expectedValue float64 = 600

	actualValue, err := FetchValue(url, path, data, contentType, regex)

	if err == nil && actualValue == expectedValue {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly. %s", err.Error())
	}

}

func Test_FetchValue_Invalid_Content_Type(t *testing.T) {

	url := ""
	path := ""
	data := []byte("")
	contentType := "invalidContentType"
	regex := ""

	_, err := FetchValue(url, path, data, contentType, regex)

	if err != nil && err == err.(ErrorParserInvalidContentType) {
		t.Log("Test succeeded as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly. %s", err.Error())
	}

}
