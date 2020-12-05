package validator

import "fmt"

var (
	errHeight = fmt.Errorf("'height' should follow: int && > 0")
	errWidth  = fmt.Errorf("'width' should follow: int && > 0")
)
