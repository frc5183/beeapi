package response

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`

	// CanBeFixed is whether this error can be fixed by the user.
	CanBeFixed bool `json:"can_be_fixed"`

	// NativeError is the native error which caused this error.
	NativeError error `json:"-"`
}

// CreateError creates an Error.
func CreateError(code ErrorCode, message string, canBeFixed bool) *Error {
	return &Error{
		Code:    code,
		Message: message,

		CanBeFixed: canBeFixed,
	}
}

// CreateNativeError creates an Error with NativeError set.
// TODO: Native errors are always internal server errors, this should probably be changed sometime in the future.
func CreateNativeError(nativeError error) *Error {
	return &Error{
		Code:    ErrorCodeInternalServerError,
		Message: "Internal Server Error.",

		CanBeFixed: false,

		NativeError: nativeError,
	}
}
