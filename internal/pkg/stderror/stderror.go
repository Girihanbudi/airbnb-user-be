package stderror

type StdError struct {
	HttpCode      int     `json:"-"`
	InternalCode  string  `json:"internalCode"`
	Message       string  `json:"message"`
	Error         error   `json:"-"`
	ErrorMessages *string `json:"error,omitempty"`
}

func New(code int, internalCode string, message string) StdError {
	return StdError{
		HttpCode:     code,
		InternalCode: internalCode,
		Message:      message,
	}
}

func (err *StdError) ErrorMsg(message error) *StdError {
	err.Error = message
	return err
}
