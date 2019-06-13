package types

import (
	"encoding/json"
	. "prometheus-metrics-exporter/pmeerrors"
)

type Metric struct {
	Name           string `json:"name"`
	Help           string `json:"help"`
	Path           string `json:"path"`
	InstrumentType string `json:"instrumentType"`
	Regex          string `json:"regex"`
}

func (m *Metric) UnmarshalJSON(data []byte) error {
	type Alias Metric
	var t Alias

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	if t.Name == "" {
		return ErrorMetricUnmarshal{Err: "Metric: Name is empty"}
	}

	if t.Help == "" {
		return ErrorMetricUnmarshal{Err: "Metric: Help is empty"}
	}

	if t.Path == "" {
		return ErrorMetricUnmarshal{Err: "Metric: Path is empty"}
	}

	if t.InstrumentType == "" {
		return ErrorMetricUnmarshal{Err: "Metric: InstrumentType is empty"}
	}

	m.Name = t.Name
	m.Help = t.Help
	m.Path = t.Path
	m.InstrumentType = t.InstrumentType
	m.Regex = t.Regex

	return nil
}
