package errors

import "net/http"

type ErrorResponse struct {
	HttpCode int
	Errors   error
}

var (
	BadRequest    *ErrorResponse = initError(http.StatusBadRequest)
	NotFound      *ErrorResponse = initError(http.StatusNotFound)
	InternalError *ErrorResponse = initError(http.StatusInternalServerError)
)

func initError(httpCode int) *ErrorResponse {
	return &ErrorResponse{HttpCode: httpCode}
}

func (EC *ErrorResponse) NewError(errors error) *ErrorResponse {
	EC.Errors = errors
	return EC
}
