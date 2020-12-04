package business

import "net/http"

// Cover : Implements the Cover service methods.
var Cover = &coverService{}

type coverService struct{}

func (c *coverService) GetRandomCover(args map[string]interface{}) (*http.Response, error) {
	return &http.Response{StatusCode: 200}, nil
}
