package response

type ErrorCode string

const (
	ErrorCodeInvalidRequest   ErrorCode = "invalid_request"
	ErrorCodeFailedValidation ErrorCode = "failed_validation"

	ErrorCodeInternalServerError ErrorCode = "internal_error"
)
