package errors

type ErrorResponse struct {
	HttpCode int
	Errors   error
}

var (
	NotFound      *ErrorResponse = InitError(404)
	InternalError *ErrorResponse = InitError(500)
)

func InitError(httpCode int) *ErrorResponse {
	return &ErrorResponse{HttpCode: httpCode}
}

func (EC *ErrorResponse) NewError(errors error) *ErrorResponse {
	EC.Errors = errors
	return EC
}
