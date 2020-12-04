package exception

import "net/http"

// Validation :
var Validation = func(message string) *Exception {
	return &Exception{http.StatusBadRequest, "VALIDATION_ERROR", or(message, "validation error")}
}

// Unexpected : Internal Server Error
var Unexpected = func(message string) *Exception {
	return &Exception{http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", or(message, "unexpected error")}
}
