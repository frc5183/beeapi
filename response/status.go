package response

type Status string

const (
	StatusSuccess Status = "success" // StatusSuccess is for successful requests.
	StatusWarning Status = "warning" // StatusWarning is for errors which are not fatal and can be ignored, however may have unintended side effects.
	StatusFatal   Status = "error"   // StatusFatal is for errors which must be fixed on the client or server side, depending on the error.
)
