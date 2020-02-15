package derr

type InValidArgumentErr struct {
	err error
}

func NewInValidArgumentErr(err error) *InValidArgumentErr {
	return &InValidArgumentErr{err: err}
}

func (e *InValidArgumentErr) Error() string {
	return e.err.Error()
}

func (e *InValidArgumentErr) Unwrap() error {
	return e.err
}
