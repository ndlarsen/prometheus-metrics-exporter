package htmlparser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	. "prometheus-metrics-exporter/internal/pmeerrors"
	"regexp"
	"strconv"
)

func FetchValue(path string, reader io.Reader, pattern string) (float64, error) {

	value, err := Extract(path, reader)

	if err != nil {
		return -1, err
	}

	value, err = match(value, pattern)
	if err != nil {
		return -1, err
	}

	f, err := strconv.ParseFloat(value, 64)

	if err != nil {
		errStr := fmt.Sprintf("HTML parsing: Unable to parse as float \"%s\"", value)
		return f, ErrorHtmlParserTypeConversion{Err: errStr}
	}

	return f, nil

}

func Extract(path string, reader io.Reader) (string, error) {
	html, err := htmlquery.Parse(reader)

	if err != nil {
		return "", ErrorHtmlParserParsing{Err: err.Error()}
	}

	resArr := htmlquery.Find(html, path)

	if len(resArr) < 1 {
		errStr := "HTML parsing: No such element found"
		return "", ErrorHtmlParserNoSuchElement{Err: errStr}
	}

	if len(resArr) > 1 {
		errStr := "HTML parsing: More than one element found"
		return "", ErrorHtmlParserTooManyElements{Err: errStr}
	}

	value := htmlquery.InnerText(resArr[0])
	return value, nil
}

func match(content string, pattern string) (string, error) {
	if pattern == "" {
		return content, nil
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", ErrorHtmlParserRegexCompileError{Err: err.Error()}
	}

	subMatchAll := r.FindStringSubmatch(content)

	if subMatchAll == nil {
		errString := fmt.Sprintf("HTML parsing: No match for regex \"%s\" found", pattern)
		return "", ErrorHtmlParserRegexNoMatch{Err: errString}
	} else if len(subMatchAll) == 1 {
		errString := "HTML parsing: No capture group supplied"
		return "", ErrorHtmlParserRegexNoCaptureGroup{Err: errString}
	} else {
		return subMatchAll[1], nil
	}
}
