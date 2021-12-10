package errs

type ApiErr struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
}

func (a ApiErr) Error() string {
	return a.Message
}
func (a ApiErr) AsMessage() *ApiErr {
	return &ApiErr{
		Message: a.Error(),
	}
}
