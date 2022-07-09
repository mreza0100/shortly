package er

import (
	"fmt"
	"net/http"
)

func Status(err error) int {
	fmt.Println(err)

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
		fallthrough
	default:
		return http.StatusInternalServerError
	}
}
