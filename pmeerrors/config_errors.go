package pmeerrors

// Config related
type ErrorConfigReadFile struct {
	Err string
}

func (c ErrorConfigReadFile) Error() string {
	return c.Err
}

type ErrorConfigConversion struct {
	Err string
}

func (c ErrorConfigConversion) Error() string {
	return c.Err
}

type ErrorConfigUnmarshal struct {
	Err string
}

func (c ErrorConfigUnmarshal) Error() string {
	return c.Err
}
