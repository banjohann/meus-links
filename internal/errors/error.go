package error

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newHttpError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}

func ErrorBadRequest(message string) *HttpError {
	return newHttpError(message, 400)
}

func ErrorUnauthorized(message string) *HttpError {
	return newHttpError(message, 401)
}

func ErrorNotFound(message string) *HttpError {
	return newHttpError(message, 404)
}

func ErrorInternal(message string) *HttpError {
	return newHttpError(message, 500)
}
