package services

type Services struct {
	User UserServicePort
}

type UserServicePort interface {
	Signup(email, password string) error
	Login(email, password string) (string, error)
}
