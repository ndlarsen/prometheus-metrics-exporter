package types

import (
	"encoding/json"
	"prometheus-metrics-exporter/internal/pmeerrors/scrapetarget"
)

type ScrapeTarget struct {
	Url           string     `json:"url"`
	BasicAuth     *BasicAuth `json:"basicAuth,omitempty"`
	Metrics       []Metric   `json:"metrics"`
	Labels        []Label    `json:"labels"`
	MimeType      string     `json:"mimeType"`
	JobName       string     `json:"jobName"`
	TimeoutInSecs int        `json:"timeoutInSecs"`
}

func (st *ScrapeTarget) UnmarshalJSON(data []byte) error {
	type Alias ScrapeTarget
	var t Alias

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	if t.Url == "" {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: Url is empty"}
	}

	if t.Metrics == nil || len(t.Metrics) < 1 {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: Metrics is empty"}
	}

	if t.Labels == nil || len(t.Labels) < 1 {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: Labels is empty"}
	}

	if t.MimeType == "" {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: MimeType is empty"}
	}

	if t.JobName == "" {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: JobName is empty"}
	}

	if t.TimeoutInSecs < 1 {
		return scrapetarget.ErrorScrapeTargetUnmarshal{Err: "ScrapeTarget: TimeoutInSecs is empty"}
	}

	if t.BasicAuth != nil {
		st.BasicAuth = t.BasicAuth
	}

	st.Url = t.Url
	st.Metrics = t.Metrics
	st.Labels = t.Labels
	st.MimeType = t.MimeType
	st.JobName = t.JobName
	st.TimeoutInSecs = t.TimeoutInSecs

	return nil
}
