package pmeerrors

// Label realted
type ErrorLabelUnmarshal struct {
	Err string
}

func (e ErrorLabelUnmarshal) Error() string {
	return e.Err
}
