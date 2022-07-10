package customerror

import (
	"net/http"
)

// Handeling status codes gracefully
func Status(err error) int {
	switch err {
	case NotFound:
		return http.StatusNotFound
	case EmailAlreadyExists:
		return http.StatusConflict
	case InvalidEmail:
		return http.StatusBadRequest
	case InvalidCredentials:
		return http.StatusUnauthorized
	case InvalidToken:
		return http.StatusBadRequest
	case ExpiredToken:
		return http.StatusUnauthorized
	case InvalidURL:
		return http.StatusBadRequest

	case GeneralFailure:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
