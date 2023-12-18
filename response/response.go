package response

// Response is the response object returned by the API.
type Response struct {
	// Status is the status of the response.
	Status Status `json:"status"`

	// Message is the message of the response.
	Message string `json:"message,omitempty"`

	// Data is the data of the response.
	// Should be empty if status is StatusFatal.
	Data interface{} `json:"data,omitempty"`

	// Errors is the errors of the response.
	// Should be empty if status is StatusSuccess.
	Errors []*Error `json:"errors,omitempty"`

	// HTTPCode is the HTTP code of the response.
	HTTPCode int `json:"-"`
}

func CreateSuccessResponse(message string, data interface{}, HTTPCode int) *Response {
	return &Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,

		HTTPCode: HTTPCode,
	}
}

func CreateWarningResponse(message string, data interface{}, errors []*Error, HTTPCode int) *Response {
	return &Response{
		Status:  StatusWarning,
		Message: message,
		Data:    data,
		Errors:  errors,

		HTTPCode: HTTPCode,
	}
}

func CreateFatalResponse(message string, errors []*Error, HTTPCode int) *Response {
	return &Response{
		Status:  StatusFatal,
		Message: message,
		Errors:  errors,

		HTTPCode: HTTPCode,
	}
}
