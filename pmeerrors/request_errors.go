package pmeerrors

// Request related
type ErrorRequestTimeOut struct {
	Err string
}

func (r ErrorRequestTimeOut) Error() string {
	return r.Err
}

type ErrorRequestResponseStatus404 struct {
	Err string
}

func (c ErrorRequestResponseStatus404) Error() string {
	return c.Err
}

type ErrorRequestResponseStatus500 struct {
	Err string
}

func (r ErrorRequestResponseStatus500) Error() string {
	return r.Err
}

type ErrorRequestResponseStatusNot200 struct {
	Err string
}

func (r ErrorRequestResponseStatusNot200) Error() string {
	return r.Err
}

type ErrorRequestInvalidContentTypeFound struct {
	Err string
}

func (r ErrorRequestInvalidContentTypeFound) Error() string {
	return r.Err
}

type ErrorRequestUnableToReadBody struct {
	Err string
}

func (r ErrorRequestUnableToReadBody) Error() string {
	return r.Err
}

type ErrorRequestContentTypeParse struct {
	Err string
}

func (r ErrorRequestContentTypeParse) Error() string {
	return r.Err
}
