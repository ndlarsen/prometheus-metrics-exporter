package jsonparser

// Json parser related
type ErrorJsonParserInvalidType struct {
	Err string
}

func (j ErrorJsonParserInvalidType) Error() string {
	return j.Err
}

type ErrorJsonParserTypeConversion struct {
	Err string
}

func (j ErrorJsonParserTypeConversion) Error() string {
	return j.Err
}

type ErrorJsonParserValueEmpty struct {
	Err string
}

func (j ErrorJsonParserValueEmpty) Error() string {
	return j.Err
}
