package scrapetarget_test

import (
	"os"
	. "prometheus-metrics-exporter/internal/pmeerrors/scrapetarget"
	. "prometheus-metrics-exporter/test"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_ErrorScrapeTargetUnmarshal_Non_empty_string(t *testing.T) {

	expectedStr := "NonEmptyString"
	err := ErrorScrapeTargetUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}

func Test_ErrorScrapeTargetUnmarshal_Empty_string(t *testing.T) {

	expectedStr := ""
	err := ErrorScrapeTargetUnmarshal{Err: expectedStr}

	Compare(t, err, expectedStr)
}
