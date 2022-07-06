package er

import "errors"

// Errors for token validation & security
var (
	InvalidToken = errors.New("Invalid token")
	ExpiredToken = errors.New("Expired token")
)
