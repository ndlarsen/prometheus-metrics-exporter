package matcher

import (
	"fmt"
	. "prometheus-metrics-exporter/internal/pmeerrors"
	"regexp"
)

func Match(content string, pattern string) (string, error) {
	if pattern == "" {
		return content, nil
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", ErrorMatcherRegexCompileError{Err: err.Error()}
	}

	subMatchAll := r.FindStringSubmatch(content)

	if subMatchAll == nil {
		errString := fmt.Sprintf("HTML parsing: No match for regex \"%s\" found", pattern)
		return "", ErrorMatcherRegexNoMatch{Err: errString}
	} else if len(subMatchAll) == 1 {
		errString := "HTML parsing: No capture group supplied"
		return "", ErrorMatcherRegexNoCaptureGroup{Err: errString}
	} else {
		return subMatchAll[1], nil
	}
}
