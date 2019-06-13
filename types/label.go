package types

import (
	"encoding/json"
	. "prometheus-metrics-exporter/pmeerrors"
)

type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (l *Label) UnmarshalJSON(data []byte) error {
	type Alias Label
	var t Alias

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	if t.Name == "" {
		return ErrorLabelUnmarshal{Err: "Label: Name is empty"}
	}

	if t.Value == "" {
		return ErrorLabelUnmarshal{Err: "Label: Value is empty"}
	}

	l.Name = t.Name
	l.Value = t.Value

	return nil
}
