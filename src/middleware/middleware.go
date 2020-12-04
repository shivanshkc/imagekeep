package middleware

import (
	"context"
	"encoding/json"
	"gold/src/logger"
	"io/ioutil"
	"net/http"
	"time"
)

var log = logger.Get()

// Interceptor : The very first method to be invoked after request reception.
func Interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		arrival := time.Now()

		log.Sugar().Infof("New Request %s %s, Timestamp: %d", req.Method, req.URL, arrival.UnixNano())
		next.ServeHTTP(writer, req)
		log.Sugar().Infof("Request took %dms to process.", time.Since(arrival).Milliseconds())
	})
}

// CORS : Handles the CORS issues.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")

		if req.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(writer, req)
	})
}

// ContentApplicationJSON : Sets the response content-type to application/json
func ContentApplicationJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("content-type", "application/json")
		next.ServeHTTP(writer, req)
	})
}

// BodyParser : Converts the io.Reader body into map[string]interface{}
func BodyParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Sugar().Warnf("Failed to read request body: %s", err.Error())
		}

		body := map[string]interface{}{}
		err = json.Unmarshal(bytes, &body)
		if err != nil {
			log.Sugar().Warnf("Failed to convert body byte array into map: %s", err.Error())
			body = nil
		}

		ctx := context.WithValue(req.Context(), BodyKey, body)
		next.ServeHTTP(writer, req.WithContext(ctx))
	})
}
