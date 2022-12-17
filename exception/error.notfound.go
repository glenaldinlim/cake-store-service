package exception

type NotFoundError struct {
	Error string
}

func NewNotFounderror(error string) NotFoundError {
	return NotFoundError{Error: error}
}
