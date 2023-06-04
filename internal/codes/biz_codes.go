package codes

import "net/http"

// http status, bisiness code, message.
var (
	CodeOK = New(http.StatusOK, "OK", "")
	//nolint: gomnd
	CodePartSuccess = New(202, "PartSuccess", "part success")

	CodePermissionDenied = New(http.StatusUnauthorized, "AuthFailed", "authentication failed")
	CodeNotAuthorized    = New(http.StatusForbidden, "NotAuthorized", "resource is not authorized")
	CodeNotFound         = New(http.StatusNotFound, "NotFound", "resource does not exist")
	CodeValidationFailed = New(http.StatusBadRequest, "ValidationFailed", "validation failed")
	CodeNotAvailable     = New(http.StatusBadRequest, "NotAvailable", "not available")

	CodeInternal = New(http.StatusInternalServerError, "InternalError", "an error occurred internally")
)
