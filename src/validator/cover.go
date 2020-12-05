package validator

import (
	"github.com/hisitra/valkyrie/v2"
)

// Cover : Implements the cover validation methods.
var Cover = &coverValidator{}

type coverValidator struct{}

func (c *coverValidator) GetRandomCover(args map[string]interface{}) error {
	return valkyrie.PureMap().
		Key("height", false, heightRule).
		Key("width", false, widthRule).
		Apply(args)
}
