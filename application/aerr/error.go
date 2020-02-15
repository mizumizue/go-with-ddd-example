package aerr

type ResourceDuplicateErr struct {
	err error
}

func NewResorseDuplicateErr(err error) *ResourceDuplicateErr {
	return &ResourceDuplicateErr{err: err}
}

func (e *ResourceDuplicateErr) Error() string {
	return e.err.Error()
}

func (e *ResourceDuplicateErr) Unwrap() error {
	return e.err
}
