package pmeparser

import (
	"bytes"
	"fmt"
	"prometheus-metrics-exporter/internal/htmlparser"
	"prometheus-metrics-exporter/internal/jsonparser"
	"prometheus-metrics-exporter/internal/pmeerrors/pmeparser"
)

func FetchValue(url string, path string, data []byte, contentType string, regex string) (float64, error) {

	if contentType == "json" {
		return jsonparser.FetchValue(path, data)
	} else if contentType == "html" {
		return htmlparser.FetchValue(path, bytes.NewBuffer(data), regex)
	} else {
		errStr := fmt.Sprintf("Parsing: Invalid content type supplied \"%s\" on \"%s\"", contentType, url)
		return -1, pmeparser.ErrorParserInvalidContentType{Err: errStr}
	}

}
