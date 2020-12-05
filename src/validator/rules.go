package validator

import "github.com/hisitra/valkyrie/v2"

var heightRule = valkyrie.StringInt().
	Allow("").
	GT(0).
	WithError(errHeight)

var widthRule = valkyrie.StringInt().
	Allow("").
	GT(0).
	WithError(errWidth)
