package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

// JSON scraping tests
func Test_check_json_gauge_no_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_no_regex_no_basic_auth_gauge",
		"test json help value no regex no basic auth gauge",
		"TestJsonNoRegexNoBasicAuthLabelNameGauge",
		"TestJsonNoRegexNoBasicAuthLabelValueGauge",
		"TestJsonNoRegexNoBasicAuthJobNameGauge",
		"gauge",
		"65",
	}

	doTest(t, false, values)

}

func Test_check_json_gauge_no_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_no_regex_basic_auth_gauge",
		"test json help value no regex basic auth gauge",
		"TestJsonNoRegexBasicAuthLabelNameGauge",
		"TestJsonNoRegexBasicAuthLabelValueGauge",
		"TestJsonNoRegexBasicAuthJobNameGauge",
		"gauge",
		"65",
	}

	doTest(t, true, values)
}

func Test_check_json_counter_no_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_no_regex_no_basic_auth_counter",
		"test json help value no regex no basic auth counter",
		"TestJsonNoRegexNoBasicAuthLabelNameCounter",
		"TestJsonNoRegexNoBasicAuthLabelValueCounter",
		"TestJsonNoRegexNoBasicAuthJobNameCounter",
		"counter",
		"65",
	}

	doTest(t, false, values)

}

func Test_check_json_counter_no_regex_basicAuth(t *testing.T) {
	values := []string{
		"test_json_help_name_no_regex_basic_auth_counter",
		"test json help value no regex basic auth counter",
		"TestJsonNoRegexBasicAuthLabelNameCounter",
		"TestJsonNoRegexBasicAuthLabelValueCounter",
		"TestJsonNoRegexBasicAuthJobNameCounter",
		"counter",
		"65",
	}

	doTest(t, true, values)
}

func Test_check_json_gauge_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_regex_no_basic_auth_gauge",
		"test json help value regex no basic auth gauge",
		"TestJsonRegexNoBasicAuthLabelNameGauge",
		"TestJsonRegexNoBasicAuthLabelValueGauge",
		"TestJsonRegexNoBasicAuthJobNameGauge",
		"gauge",
		"996",
	}

	doTest(t, false, values)

}

func Test_check_json_gauge_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_regex_basic_auth_gauge",
		"test json help value regex basic auth gauge",
		"TestJsonRegexBasicAuthLabelNameGauge",
		"TestJsonRegexBasicAuthLabelValueGauge",
		"TestJsonRegexBasicAuthJobNameGauge",
		"gauge",
		"996",
	}

	doTest(t, true, values)
}

func Test_check_json_counter_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_regex_no_basic_auth_counter",
		"test json help value regex no basic auth counter",
		"TestJsonRegexNoBasicAuthLabelNameCounter",
		"TestJsonRegexNoBasicAuthLabelValueCounter",
		"TestJsonRegexNoBasicAuthJobNameCounter",
		"counter",
		"996",
	}

	doTest(t, false, values)

}

func Test_check_json_counter_regex_basicAuth(t *testing.T) {
	values := []string{
		"test_json_help_name_regex_basic_auth_counter",
		"test json help value regex basic auth counter",
		"TestJsonRegexBasicAuthLabelNameCounter",
		"TestJsonRegexBasicAuthLabelValueCounter",
		"TestJsonRegexBasicAuthJobNameCounter",
		"counter",
		"996",
	}

	doTest(t, true, values)
}

// HTML scraping tests
func Test_check_html_gauge_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_no_basic_auth_gauge",
		"test html help value no basic auth gauge",
		"TestHtmlNoBasicAuthLabelNameGauge",
		"TestHtmlNoBasicAuthLabelValueGauge",
		"TestHtmlNoBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTest(t, false, values)
}

func Test_check_html_gauge_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_basic_auth_gauge",
		"test html help value basic auth gauge",
		"TestHtmlBasicAuthLabelNameGauge",
		"TestHtmlBasicAuthLabelValueGauge",
		"TestHtmlBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTest(t, true, values)
}

func Test_check_html_counter_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_no_basic_auth_counter",
		"test html help value no basic auth counter",
		"TestHtmlNoBasicAuthLabelNameCounter",
		"TestHtmlNoBasicAuthLabelValueCounter",
		"TestHtmlNoBasicAuthJobNameCounter",
		"counter",
		"567",
	}

	doTest(t, false, values)
}

func Test_check_html_counter_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_basic_auth_counter",
		"test html help value basic auth counter",
		"TestHtmlBasicAuthLabelNameCounter",
		"TestHtmlBasicAuthLabelValueCounter",
		"TestHtmlBasicAuthJobNameCounter",
		"counter",
		"567",
	}

	doTest(t, true, values)
}

func doTest(t *testing.T, withBasicAuth bool, values []string) {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	url := "http://pushgateway:9091/metrics"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		errStr := fmt.Sprintf("http client failed: %s", err)
		t.Fatal(errStr)
	}

	if withBasicAuth {
		req.SetBasicAuth("username", "password")
	}

	response, err := client.Do(req)

	if err != nil {
		errStr := fmt.Sprintf("http client failed: %s", err)
		t.Fatal(errStr)
	}

	body, err := ioutil.ReadAll(response.Body)

	defer func() {
		if _err := response.Body.Close(); _err != nil {
			err = _err
		}
	}()

	if err != nil {
		errStr := fmt.Sprintf("reading response body failed: %s", err)
		t.Fatal(errStr)
	}

	bodyStr := string(body)

	str1 := fmt.Sprintf("# HELP %s %s", values[0], values[1])
	str2 := fmt.Sprintf("# TYPE %s %s", values[0], values[5])
	str3 := fmt.Sprintf("%s{%s=\"%s\",hostname=\"simplewebserver\",instance=\"\",job=\"%s\"} %s", values[0], values[2], values[3], values[4], values[6])

	stringArr := []string{str1, str2, str3}

	for _, s := range stringArr {
		if !strings.Contains(bodyStr, s) {
			t.Errorf("Expected string: \"%s\" was not found", s)
		}
	}

}
