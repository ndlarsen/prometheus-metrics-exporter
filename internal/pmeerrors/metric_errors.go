package pmeerrors

//Metrics related
type ErrorMetricUnmarshal struct {
	Err string
}

func (e ErrorMetricUnmarshal) Error() string {
	return e.Err
}
