package apierrors

type APIError struct {
	ErrCode ErrCode `json:"err_code"`
	Message string  `json:"message"`
	Err     error   `json:"-"`
}

func (myErr *APIError) Error() string {
	return myErr.Err.Error()
}

func (myErr *APIError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &APIError{ErrCode: code, Message: message, Err: err}
}
