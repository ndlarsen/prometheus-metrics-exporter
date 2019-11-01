package instrument

// instrument related
type ErrorInstrumentMissingValue struct {
	Err string
}

func (i ErrorInstrumentMissingValue) Error() string {
	return i.Err
}

type ErrorInstrumentUnsupportedType struct {
	Err string
}

func (i ErrorInstrumentUnsupportedType) Error() string {
	return i.Err
}

type ErrorInstrumentUrlParse struct {
	Err string
}

func (i ErrorInstrumentUrlParse) Error() string {
	return i.Err
}

type ErrorInstrumentPushFailed struct {
	Err string
}

func (i ErrorInstrumentPushFailed) Error() string {
	return i.Err
}
