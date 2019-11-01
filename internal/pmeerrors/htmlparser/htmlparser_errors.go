package htmlparser

// html parser related
type ErrorHtmlParserTypeConversion struct {
	Err string
}

func (h ErrorHtmlParserTypeConversion) Error() string {
	return h.Err
}

type ErrorHtmlParserParsing struct {
	Err string
}

func (h ErrorHtmlParserParsing) Error() string {
	return h.Err
}

type ErrorHtmlParserNoSuchElement struct {
	Err string
}

func (h ErrorHtmlParserNoSuchElement) Error() string {
	return h.Err
}

type ErrorHtmlParserTooManyElements struct {
	Err string
}

func (h ErrorHtmlParserTooManyElements) Error() string {
	return h.Err
}
