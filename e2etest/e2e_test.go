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

	doTestOk(t, false, values, true)

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

	doTestOk(t, true, values, true)
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

	doTestOk(t, false, values, true)

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

	doTestOk(t, true, values, true)
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

	doTestOk(t, false, values, true)

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

	doTestOk(t, true, values, true)
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

	doTestOk(t, false, values, true)

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

	doTestOk(t, true, values, true)
}

// HTML scraping tests
func Test_check_html_gauge_no_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_no_regex_no_basic_auth_gauge",
		"test html help value no regex no basic auth gauge",
		"TestHtmlNoRegexNoBasicAuthLabelNameGauge",
		"TestHtmlNoRegexNoBasicAuthLabelValueGauge",
		"TestHtmlNoRegexNoBasicAuthJobNameGauge",
		"gauge",
		"223",
	}

	doTestOk(t, false, values, true)
}

func Test_check_html_gauge_no_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_name_no_regex_basic_auth_gauge",
		"test html help value no regex basic auth gauge",
		"TestHtmlNoRegexBasicAuthLabelNameGauge",
		"TestHtmlNoRegexBasicAuthLabelValueGauge",
		"TestHtmlNoRegexBasicAuthJobNameGauge",
		"gauge",
		"223",
	}

	doTestOk(t, true, values, true)
}

func Test_check_html_counter_no_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_no_regex_name_no_basic_auth_counter",
		"test html help no regex value no basic auth counter",
		"TestHtmlNoRegexNoBasicAuthLabelNameCounter",
		"TestHtmlNoRegexNoBasicAuthLabelValueCounter",
		"TestHtmlNoRegexNoBasicAuthJobNameCounter",
		"counter",
		"223",
	}

	doTestOk(t, false, values, true)
}

func Test_check_html_counter_no_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_no_regex_name_basic_auth_counter",
		"test html help no regex value basic auth counter",
		"TestHtmlNoRegexBasicAuthLabelNameCounter",
		"TestHtmlNoRegexBasicAuthLabelValueCounter",
		"TestHtmlNoRegexBasicAuthJobNameCounter",
		"counter",
		"223",
	}

	doTestOk(t, true, values, true)
}

func Test_check_html_gauge_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_regex_name_no_basic_auth_gauge",
		"test html help regex value no basic auth gauge",
		"TestHtmlRegexNoBasicAuthLabelNameGauge",
		"TestHtmlRegexNoBasicAuthLabelValueGauge",
		"TestHtmlRegexNoBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTestOk(t, false, values, true)
}

func Test_check_html_gauge_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_regex_name_basic_auth_gauge",
		"test html help regex value basic auth gauge",
		"TestHtmlRegexBasicAuthLabelNameGauge",
		"TestHtmlRegexBasicAuthLabelValueGauge",
		"TestHtmlRegexBasicAuthJobNameGauge",
		"gauge",
		"567",
	}

	doTestOk(t, true, values, true)
}

func Test_check_html_counter_regex_no_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_regex_name_no_basic_auth_counter",
		"test html help regex value no basic auth counter",
		"TestHtmlRegexNoBasicAuthLabelNameCounter",
		"TestHtmlRegexNoBasicAuthLabelValueCounter",
		"TestHtmlRegexNoBasicAuthJobNameCounter",
		"counter",
		"567",
	}

	doTestOk(t, false, values, true)
}

func Test_check_html_counter_regex_basicAuth(t *testing.T) {

	values := []string{
		"test_html_help_regex_name_basic_auth_counter",
		"test html help regex value basic auth counter",
		"TestHtmlRegexBasicAuthLabelNameCounter",
		"TestHtmlRegexBasicAuthLabelValueCounter",
		"TestHtmlRegexBasicAuthJobNameCounter",
		"counter",
		"567",
	}

	doTestOk(t, true, values, true)
}

func Test_check_html_basicAuth_401_unauthorized_no_credentials(t *testing.T) {

	values := []string{
		"test_html_basic_auth_job_name_401_unauthorized_no_credentials",
		"test html basic auth job name 401 unauthorized no credentials",
		"TestHtmlBasicAuthJobName401UnauthorizedNoCredentialsLabelName",
		"TestHtmlBasicAuthJobName401UnauthorizedNoCredentialsLabelValue",
		"TestHtmlBasicAuthJobName401UnauthorizedNoCredentials",
		"gauge",
		"567",
	}

	doTestOk(t, true, values, false)
}

func Test_check_jon_basicAuth_401_unauthorized_no_credentials(t *testing.T) {

	values := []string{
		"test_json_basic_auth_job_name_401_unauthorized_no_credentials",
		"test json basic auth job name 401 unauthorized no credentials",
		"TestJsonBasicAuthJobName401UnauthorizedNoCredentialsLabelName",
		"TestJsonBasicAuthJobName401UnauthorizedNoCredentialsLabelValue",
		"TestJsonBasicAuthJobName401UnauthorizedNoCredentials",
		"gauge",
		"567",
	}

	doTestOk(t, true, values, false)
}

func Test_check_html_basicAuth_401_unauthorized_invalid_credentials(t *testing.T) {

	values := []string{
		"test_html_basic_auth_job_name_401_unauthorized_invalid_credentials",
		"test html basic auth job name 401 unauthorized invalid credentials",
		"TestHtmlBasicAuthJobName401UnauthorizedInvalidCredentialsLabelName",
		"TestHtmlBasicAuthJobName401UnauthorizedInvalidCredentialsLabelValue",
		"TestHtmlBasicAuthJobName401UnauthorizedInvalidCredentials",
		"gauge",
		"567",
	}

	doTestOk(t, true, values, false)
}

func Test_check_jon_basicAuth_401_unauthorized_invalid_credentials(t *testing.T) {

	values := []string{
		"test_json_basic_auth_job_name_401_unauthorized_no_credentials",
		"test json basic auth job name 401 unauthorized no credentials",
		"TestJsonBasicAuthJobName401UnauthorizedInvalidCredentialsLabelName",
		"TestJsonBasicAuthJobName401UnauthorizedInvalidCredentialsLabelValue",
		"TestJsonBasicAuthJobName401UnauthorizedInvalidCredentials",
		"gauge",
		"567",
	}

	doTestOk(t, true, values, false)
}

func doTestOk(t *testing.T, withBasicAuth bool, values []string, valuesPresent bool) {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	url := "http://pushgateway:9091/metrics"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		errStr := fmt.Sprintf("http client failed: %s", err)
		t.Fatal(errStr)
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
		if valuesPresent && !strings.Contains(bodyStr, s) {
			t.Errorf("String: \"%s\" was not found", s)
		} else if !valuesPresent && strings.Contains(bodyStr, s) {
			t.Errorf("String: \"%s\" was found", s)
		}
	}

}
