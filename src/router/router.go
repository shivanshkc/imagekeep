package router

import (
	"fmt"
	"gold/src/configs"
	"gold/src/exception"
	"gold/src/logger"
	"gold/src/middleware"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var conf = configs.Get()
var log = logger.Get()

// Get : Gets all the HTTP Handlers to be used in http.ListenAndServe.
func Get() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.Interceptor)
	router.Use(middleware.CORS)

	apiRouter := attachAPI(router.PathPrefix("/api").Subrouter())
	_ = attachGeneralAccess(apiRouter.PathPrefix("/generalAccess").Subrouter())

	return router
}

func attachAPI(router *mux.Router) *mux.Router {
	router.Use(middleware.ContentApplicationJSON)
	router.Use(middleware.BodyParser)

	router.HandleFunc("", func(writer http.ResponseWriter, req *http.Request) {
		resJSON := fmt.Sprintf(
			`{"name":"%s","version":"%s"}`,
			conf.Application.Name,
			conf.Application.Version,
		)
		_, _ = writer.Write([]byte(resJSON))
	}).Methods(http.MethodGet, http.MethodOptions)

	return router
}

func baseHandler(
	writer http.ResponseWriter,
	args map[string]interface{},
	validator func(args map[string]interface{}) error,
	service func(args map[string]interface{}) (*http.Response, error),
) {
	err := validator(args)
	if err != nil {
		exception.Send(exception.Validation(err.Error()), writer)
		return
	}

	result, err := service(args)
	if err != nil {
		exception.Send(err, writer)
		return
	}

	writeResponse(writer, result)
}

func writeResponse(writer http.ResponseWriter, result *http.Response) {
	writer.WriteHeader(result.StatusCode)

	for key, value := range result.Header {
		if len(value) == 0 {
			continue
		}
		writer.Header().Set(key, value[0])
	}

	if result.Body == nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Sugar().Warnf("Failed to read response body before sending: %s", err.Error())
		return
	}

	_, err = writer.Write(bodyBytes)
	if err != nil {
		log.Sugar().Warnf("Failed to write response body: %s", err.Error())
	}
}
