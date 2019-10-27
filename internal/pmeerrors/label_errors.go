package pmeerrors

// Label related
type ErrorLabelUnmarshal struct {
	Err string
}

func (e ErrorLabelUnmarshal) Error() string {
	return e.Err
}
