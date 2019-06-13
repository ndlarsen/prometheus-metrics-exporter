package instrument

import (
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http"
	"net/http/httptest"
	. "prometheus-metrics-exporter/pmeerrors"
	"testing"
)

const (
	name                   = "metricName"
	instrumentType         = "gauge"
	path                   = "json.path"
	help                   = "helpText"
	value          float64 = 123
)

func Test_CreateInstrument_OK_Gauge(t *testing.T) {

	_, err := CreateInstrument("gauge", path, name, help, value)

	if err != nil {
		t.Fatalf("Test failed unexpectedly %s", err)
	}

}

func Test_CreateInstrument_OK_Counter(t *testing.T) {

	_, err := CreateInstrument("counter", path, name, help, value)

	if err != nil {
		t.Fatalf("Test failed unexpectedly %s", err)
	}

}

func Test_CreateInstrument_Error_Unsupported_Instrument(t *testing.T) {

	_, err := CreateInstrument("qwe", path, name, help, value)

	if err != nil && err == err.(ErrorInstrumentUnsupportedType) {
		t.Log("Parse error as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
}

func Test_CreateInstrument_Error_Missing_Required_InstrumentType(t *testing.T) {

	_, err := CreateInstrument("", path, name, help, value)

	if err != nil && err == err.(ErrorInstrumentMissingValue) {
		t.Log("Parse error as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
}

func Test_CreateInstrument_Error_Missing_Required_JsonPath(t *testing.T) {

	_, err := CreateInstrument(instrumentType, "", name, help, value)

	if err != nil && err == err.(ErrorInstrumentMissingValue) {
		t.Log("Parse error as expected.")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
}

func Test_Push_Response_Error(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	registry := push.New("localhost:1234", "testJob")

	i, err := CreateInstrument(instrumentType, path, name, help, value)

	if err != nil {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
	registry.Collector(i)

	err = Push(ts.URL, registry)
	if err != nil && err == err.(ErrorInstrumentPushFailed) {
		t.Log("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}

}

func Test_Push_Response_Url_Parse_Failure(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	registry := push.New("localhost:1234", "testJob")

	i, err := CreateInstrument(instrumentType, path, name, help, value)

	if err != nil {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
	registry.Collector(i)

	err = Push("local host", registry)

	if err != nil && err == err.(ErrorInstrumentUrlParse) {
		t.Log("Test succeeded.")
		t.Log(err)
	} else if err != nil {
		t.Fatalf("Unexpected error")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}

}

func Test_Push_OK(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	registry := push.New(ts.URL, "testJob")

	i, err := CreateInstrument(instrumentType, path, name, help, value)

	if err != nil {
		t.Fatalf("Test failed unexpectedly %s", err)
	}
	registry.Collector(i)

	err = Push(ts.URL, registry)

	if err == nil {
		t.Log("Test succeeded.")
	} else {
		t.Fatalf("Test failed unexpectedly %s", err)
	}

}
