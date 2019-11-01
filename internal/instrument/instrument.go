package instrument

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/url"
	"prometheus-metrics-exporter/internal/pmeerrors/instrument"
)

func CreateInstrument(instrumentType string, path string, name string, help string, value float64) (prometheus.Collector, error) {
	if path == "" {
		return nil, instrument.ErrorInstrumentMissingValue{Err: "Instrument creation: required field path is empty"}
	} else if instrumentType == "" {
		return nil, instrument.ErrorInstrumentMissingValue{Err: "Instrument creation: required field instrumentType is empty"}
	}

	if instrumentType == "gauge" {
		return createGauge(name, help, value), nil
	} else if instrumentType == "counter" {
		return createCounter(name, help, value), nil
	} else {
		errStr := fmt.Sprintf("Instrument creation: unsupported instrument type \"%s\"", instrumentType)
		return nil, instrument.ErrorInstrumentUnsupportedType{Err: errStr}
	}

}

func createGauge(name string, help string, value float64) prometheus.Collector {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
	g.Set(value)
	return g
}

func createCounter(name string, help string, value float64) prometheus.Collector {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
	c.Add(value)
	return c
}

func Push(scrapeUrl string, registry *push.Pusher) error {

	u, err := url.ParseRequestURI(scrapeUrl)
	if err != nil {
		errStr := fmt.Sprintf("Instrument creation: Unable to parse url \"%s\"", scrapeUrl)
		return instrument.ErrorInstrumentUrlParse{Err: errStr}
	}

	registry.Grouping("hostname", u.Hostname())

	err = registry.Push()

	if err != nil {
		errStr := fmt.Sprintf("Instrument creation: pushing failed \"%s\"", err.Error())
		return instrument.ErrorInstrumentPushFailed{Err: errStr}
	}

	return nil
}
