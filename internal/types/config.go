package types

import (
	"encoding/json"
	"prometheus-metrics-exporter/internal/pmeerrors/config"
)

type Config struct {
	ScrapeTargets  []ScrapeTarget `json:"scrapeTargets"`
	PushGatewayUrl string         `json:"pushGatewayUrl"`
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type Alias Config
	var t Alias

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	if t.PushGatewayUrl == "" {
		return config.ErrorConfigUnmarshal{Err: "Config: PushGateway url is empty"}
	}

	if t.ScrapeTargets == nil || len(t.ScrapeTargets) < 1 {
		return config.ErrorConfigUnmarshal{Err: "Config: ScrapeTarget metrics is empty"}
	}

	c.ScrapeTargets = t.ScrapeTargets
	c.PushGatewayUrl = t.PushGatewayUrl

	return nil
}
