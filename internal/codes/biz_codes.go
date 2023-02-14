package codes

//  http status, bisiness code, message
var (
	CodeOK          = New(200, "OK", "")
	CodePartSuccess = New(202, "PartSuccess", "part success")

	CodePermissionDenied = New(401, "AuthFailed", "authentication failed")
	CodeNotAuthorized    = New(403, "NotAuthorized", "resource is not authorized")
	CodeNotFound         = New(404, "NotFound", "resource does not exist")
	CodeValidationFailed = New(400, "ValidationFailed", "validation failed")
	CodeNotAvailable     = New(400, "NotAvailable", "not available")

	CodeInternal = New(500, "InternalError", "an error occurred internally")
	CodeUnknown  = New(500, "UnknownError", "unknown error")
)
