package htmlparser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	. "prometheus-metrics-exporter/internal/matcher"
	"prometheus-metrics-exporter/internal/pmeerrors/htmlparser"
	"strconv"
	"strings"
)

func FetchValue(path string, reader io.Reader, pattern string) (float64, error) {

	value, err := Extract(path, reader)

	if err != nil {
		return -1, err
	}

	if pattern != "" {
		value, err = Match(value, pattern)
		if err != nil {
			return -1, err
		}
	} else {
		value = strings.TrimSpace(value)
	}

	f, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errStr := fmt.Sprintf("HTML parsing: Unable to parse as float \"%s\"", value)
		return f, htmlparser.ErrorHtmlParserTypeConversion{Err: errStr}
	}

	return f, nil

}

func Extract(path string, reader io.Reader) (string, error) {
	html, err := htmlquery.Parse(reader)

	if err != nil {
		return "", htmlparser.ErrorHtmlParserParsing{Err: err.Error()}
	}

	resArr := htmlquery.Find(html, path)

	if len(resArr) < 1 {
		errStr := "HTML parsing: No such element found"
		return "", htmlparser.ErrorHtmlParserNoSuchElement{Err: errStr}
	}

	if len(resArr) > 1 {
		errStr := "HTML parsing: More than one element found"
		return "", htmlparser.ErrorHtmlParserTooManyElements{Err: errStr}
	}

	value := htmlquery.InnerText(resArr[0])
	return value, nil
}
