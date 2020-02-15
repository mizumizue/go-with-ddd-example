package aerr

type ResourceDuplicateErr struct {
	err error
}

func NewResourceDuplicateErr(err error) ResourceDuplicateErr {
	return ResourceDuplicateErr{err: err}
}

func (e ResourceDuplicateErr) Error() string {
	return e.err.Error()
}

func (e ResourceDuplicateErr) Unwrap() error {
	return e.err
}

type ResourceNotFoundErr struct {
	err error
}

func NewResourceNotFoundErr(err error) ResourceNotFoundErr {
	return ResourceNotFoundErr{err: err}
}

func (e ResourceNotFoundErr) Error() string {
	return e.err.Error()
}

func (e ResourceNotFoundErr) Unwrap() error {
	return e.err
}

type DatabaseErr struct {
	err error
}

func NewDatabaseErr(err error) DatabaseErr {
	return DatabaseErr{err: err}
}

func (e DatabaseErr) Error() string {
	return e.err.Error()
}

func (e DatabaseErr) Unwrap() error {
	return e.err
}
