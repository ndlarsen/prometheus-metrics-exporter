package types

import (
	"encoding/json"
	. "prometheus-metrics-exporter/pmeerrors"
)

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ba *BasicAuth) UnmarshalJSON(data []byte) error {
	type Alias BasicAuth
	var t Alias

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	if t.Username == "" {
		return ErrorBasicAuthUnmarshal{Err: "Username is empty"}
	}

	if t.Password == "" {
		return ErrorBasicAuthUnmarshal{Err: "Password is empty"}
	}

	ba.Username = t.Username
	ba.Password = t.Password

	return nil
}
