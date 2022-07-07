package er

import "errors"

// Errors for token validation & security
var (
	InvalidToken = errors.New("Invalid token")
	ExpiredToken = errors.New("Expired token")
)

// Error for wrong email or password
var (
	InvalidEmailOrPassword = errors.New("Invalid email or password")
	UsernameAlreadyExists  = errors.New("Username already exists")
)
