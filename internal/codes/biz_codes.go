package codes

//  http status, bisiness code, message
var (
	CodeOK          = New(200, "OK", "")
	CodePartSuccess = New(202, "PartSuccess", "part success")

	CodeNotAuthorized    = New(401, "NotAuthorized", "resource is not authorized")
	CodePermissionDenied = New(403, "PermissionDenied", "permission denied")
	CodeNotFound         = New(404, "NotFound", "resource does not exist")
	CodeValidationFailed = New(400, "ValidationFailed", "validation failed")
	CodeNotAvailable     = New(400, "NotAvailable", "not available")

	CodeInternal = New(500, "InternalError", "an error occurred internally")
	CodeUnknown  = New(500, "UnknownError", "unknown error")
)
