package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func attachGeneralAccess(router *mux.Router) *mux.Router {
	router.HandleFunc("/cover/random.jpg", getRandomCoverHandler).Methods(http.MethodGet, http.MethodOptions)
	return router
}

func getRandomCoverHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"height": req.URL.Query().Get("height"),
		"width":  req.URL.Query().Get("width"),
	}

	baseHandler(writer, args, coverValidator.GetRandomCover, coverService.GetRandomCover)
}
