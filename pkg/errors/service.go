package er

import "errors"

var (
	URLNotValid   = errors.New("URL is not valid")
	EmailNotValid = errors.New("Email is not valid")
)
