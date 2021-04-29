package errorException

import (
	"Backend-Loans/domain/dto"
)

type ErrorResponse interface {
	SetError(string, dto.Response)
	Error() string
	Response() dto.Response
}
