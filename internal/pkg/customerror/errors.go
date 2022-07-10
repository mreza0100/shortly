package customerror

import "errors"

// General
var (
	GeneralFailure = errors.New("General failure")
	NotFound       = errors.New("Not found")
)

// Errors for validation
var (
	InvalidToken = errors.New("Invalid token")
	ExpiredToken = errors.New("Expired token")

	InvalidURL   = errors.New("URL is not valid")
	InvalidEmail = errors.New("Email is not valid")
)

// Error for wrong email or password
var (
	InvalidCredentials = errors.New("Invalid email or password")
	EmailAlreadyExists = errors.New("Email already exists")
)
