package errorException

import (
	"Backend-Loans/domain/dto"
)

type ErrorResponseImpl struct {
	errorResponse string
	response      dto.Response
}

var _ ErrorResponse = &ErrorResponseImpl{}

func InitErrorResponse() *ErrorResponseImpl {
	return &ErrorResponseImpl{}
}

func (e *ErrorResponseImpl) SetError(err string, response dto.Response) {
	e.errorResponse = err
	e.response = response
}

func (e *ErrorResponseImpl) Error() string {
	return e.errorResponse
}

func (e *ErrorResponseImpl) Response() dto.Response {
	return e.response
}
