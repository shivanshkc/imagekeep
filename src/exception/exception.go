package exception

import (
	"fmt"
	"gold/src/logger"
	"net/http"
)

var log = logger.Get()

// Exception : The custom error type, implements the error interface.
type Exception struct {
	StatusCode int    `json:"statusCode"`
	CustomCode string `json:"customCode"`
	Message    string `json:"message"`
}

func (e *Exception) Error() string {
	return e.Message
}

// ToJSON : Converts the Exception to JSON.
func (e *Exception) ToJSON() []byte {
	return []byte(fmt.Sprintf(
		`{"statusCode":%d,"customCode":"%s","message":"%s"}`,
		e.StatusCode,
		e.CustomCode,
		e.Message,
	))
}

// Send : Sends an error safely as an HTTP response.
func Send(err error, writer http.ResponseWriter) {
	exc, ok := err.(*Exception)
	if !ok {
		log.Sugar().Errorf("Unexpected error: %s", err.Error())
		exc = Unexpected("")
	}

	writer.WriteHeader(exc.StatusCode)
	_, wErr := writer.Write(exc.ToJSON())
	if wErr != nil {
		log.Sugar().Warnf("Failed to write response in Exception Send: %s", err.Error())
	}
}
