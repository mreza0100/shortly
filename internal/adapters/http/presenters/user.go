package presenters

// Signup Endpoint
type (
	SignupRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	SignupResponse struct {
		Error string `json:"error"`
	}
)

// Signin Endpoint
type (
	SigninRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	SigninResponse struct {
		Token string `json:"token"`
		Error string `json:"error"`
	}
)
