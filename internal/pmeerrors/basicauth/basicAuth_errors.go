package basicauth

// BasicAuth related
type ErrorBasicAuthUnmarshal struct {
	Err string
}

func (e ErrorBasicAuthUnmarshal) Error() string {
	return e.Err
}
