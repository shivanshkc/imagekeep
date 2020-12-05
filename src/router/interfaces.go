package router

import (
	"gold/src/business"
	"gold/src/validator"
	"net/http"
)

var coverValidator iCoverValidator = validator.Cover
var coverService iCoverService = business.Cover

type iCoverValidator interface {
	GetRandomCover(args map[string]interface{}) error
}

type iCoverService interface {
	GetRandomCover(args map[string]interface{}) (*http.Response, error)
}
