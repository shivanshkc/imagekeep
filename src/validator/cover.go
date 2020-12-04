package validator

import (
	"fmt"
	"github.com/hisitra/valkyrie/v2"
)

// Cover : Implements the cover validation methods.
var Cover = &coverValidator{}

type coverValidator struct{}

func (c *coverValidator) GetRandomCover(args map[string]interface{}) error {
	fmt.Println("Args:", args)
	return valkyrie.PureMap().
		Key("height", false, heightRule).
		Key("width", false, widthRule).
		Apply(args)
}
