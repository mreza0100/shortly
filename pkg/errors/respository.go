package er

import "errors"

var (
	NotFound       = errors.New("Not found")
	GeneralFailure = errors.New("General failure")
)
