package pmeerrors

// pme parser related
type ErrorParserInvalidContentType struct {
	Err string
}

func (j ErrorParserInvalidContentType) Error() string {
	return j.Err
}
