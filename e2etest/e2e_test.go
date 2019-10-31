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

func Test_check_json_gauge_without_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_without_basic_auth_gauge",
		"test json help value without basic auth gauge",
		"TestJsonWithoutBasicAuthLabelNameGauge",
		"TestJsonWithoutBasicAuthLabelValueGauge",
		"TestJsonWithoutBasicAuthJobNameGauge",
		"gauge",
		"65",
	}

	doTest(t, false, values)

}

func Test_check_json_gauge_with_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_with_basic_auth_gauge",
		"test json help value with basic auth gauge",
		"TestJsonWithBasicAuthLabelNameGauge",
		"TestJsonWithBasicAuthLabelValueGauge",
		"TestJsonWithBasicAuthJobNameGauge",
		"gauge",
		"65",
	}

	doTest(t, true, values)
}

func Test_check_html_gauge_without_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_without_basic_auth_gauge",
		"test html help value without basic auth gauge",
		"TestHtmlWithoutBasicAuthLabelNameGauge",
		"TestHtmlWithoutBasicAuthLabelValueGauge",
		"TestHtmlWithoutBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTest(t, false, values)
}

func Test_check_html_gauge_with_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_with_basic_auth_gauge",
		"test html help value with basic auth gauge",
		"TestHtmlWithBasicAuthLabelNameGauge",
		"TestHtmlWithBasicAuthLabelValueGauge",
		"TestHtmlWithBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTest(t, true, values)
}

func Test_check_json_counter_without_basicAuth(t *testing.T) {

	values := []string{
		"test_json_help_name_without_basic_auth_counter",
		"test json help value without basic auth counter",
		"TestJsonWithoutBasicAuthLabelNameCounter",
		"TestJsonWithoutBasicAuthLabelValueCounter",
		"TestJsonWithoutBasicAuthJobNameCounter",
		"counter",
		"65",
	}

	doTest(t, false, values)

}

func Test_check_json_counter_with_basicAuth(t *testing.T) {
	values := []string{
		"test_json_help_name_with_basic_auth_counter",
		"test json help value with basic auth counter",
		"TestJsonWithBasicAuthLabelNameCounter",
		"TestJsonWithBasicAuthLabelValueCounter",
		"TestJsonWithBasicAuthJobNameCounter",
		"counter",
		"65",
	}

	doTest(t, true, values)
}

func Test_check_html_counter_without_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_without_basic_auth_counter",
		"test html help value without basic auth counter",
		"TestHtmlWithoutBasicAuthLabelNameCounter",
		"TestHtmlWithoutBasicAuthLabelValueCounter",
		"TestHtmlWithoutBasicAuthJobNameCounter",
		"counter",
		"567",
	}

	doTest(t, false, values)
}

func Test_check_html_counter_with_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_with_basic_auth_counter",
		"test html help value with basic auth counter",
		"TestHtmlWithBasicAuthLabelNameCounter",
		"TestHtmlWithBasicAuthLabelValueCounter",
		"TestHtmlWithBasicAuthJobNameCounter",
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

	if !(strings.Contains(bodyStr, str1) && strings.Contains(bodyStr, str2) && strings.Contains(bodyStr, str3)) {
		t.Fatal()
	}

}
