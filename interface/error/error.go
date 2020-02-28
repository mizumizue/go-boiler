package error

type ValidationErr struct {
	Err error
}

func NewValidationErr(err error) ValidationErr {
	return ValidationErr{Err: err}
}

func (e ValidationErr) Error() string {
	return e.Err.Error()
}

func (e ValidationErr) Unwrap() error {
	return e.Err
}
