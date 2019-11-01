package matcher

type ErrorMatcherRegexCompileError struct {
	Err string
}

func (h ErrorMatcherRegexCompileError) Error() string {
	return h.Err
}

type ErrorMatcherRegexNoMatch struct {
	Err string
}

func (h ErrorMatcherRegexNoMatch) Error() string {
	return h.Err
}

type ErrorMatcherRegexNoCaptureGroup struct {
	Err string
}

func (h ErrorMatcherRegexNoCaptureGroup) Error() string {
	return h.Err
}
