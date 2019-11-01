package matcher

import (
	"fmt"
	"prometheus-metrics-exporter/internal/pmeerrors/matcher"
	"regexp"
)

func Match(content string, pattern string) (string, error) {
	if pattern == "" {
		return content, nil
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", matcher.ErrorMatcherRegexCompileError{Err: err.Error()}
	}

	subMatchAll := r.FindStringSubmatch(content)

	if subMatchAll == nil {
		errString := fmt.Sprintf("Matcher: No match for regex \"%s\" found", pattern)
		return "", matcher.ErrorMatcherRegexNoMatch{Err: errString}
	} else if len(subMatchAll) == 1 {
		errString := "Matcher: No capture group supplied"
		return "", matcher.ErrorMatcherRegexNoCaptureGroup{Err: errString}
	} else {
		return subMatchAll[1], nil
	}
}
