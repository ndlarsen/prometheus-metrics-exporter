package jsonparser

import (
	"fmt"
	"github.com/tidwall/gjson"
	"prometheus-metrics-exporter/internal/pmeerrors/jsonparser"
	"strconv"
)

func FetchValue(path string, json []byte) (float64, error) {
	result := gjson.GetBytes(json, path)

	if !result.Exists() {
		errStr := fmt.Sprintf("Json parsing: no value found for \"%s\"", path)
		return -1, jsonparser.ErrorJsonParserValueEmpty{Err: errStr}
	} else if result.Type == gjson.Number {
		return result.Num, nil
	} else if result.Type == gjson.String {
		f, err := strconv.ParseFloat(result.Str, 64)

		if err != nil {
			errStr := fmt.Sprintf("Json parsing: unable to parse as float \"%s\"", result.Raw)
			return f, jsonparser.ErrorJsonParserTypeConversion{Err: errStr}
		}

		return f, nil
	} else {
		errStr := "Json parsing: invalid type found"
		return -1, jsonparser.ErrorJsonParserInvalidType{Err: errStr}
	}

}
