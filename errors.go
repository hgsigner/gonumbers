package gonumbers

import "errors"

var (
	notIntegerError     = errors.New("The value should an integer.")
	unknownBinaryPrefix = errors.New("Prefix must be binary or si.")
)
