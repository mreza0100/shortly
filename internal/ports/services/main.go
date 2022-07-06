package services

type Services struct {
	User UserServicePort
}

type UserServicePort interface {
	Signup(email, password string) error
	Signin(email, password string) (string, error)
}
