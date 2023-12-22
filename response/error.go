package response

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`

	// NativeError is the native error which caused this error.
	NativeError error `json:"-"`
}

// CreateError creates an Error.
func CreateError(code ErrorCode, message string, nativeError error) *Error {
	return &Error{
		Code:    code,
		Message: message,

		NativeError: nativeError,
	}
}
