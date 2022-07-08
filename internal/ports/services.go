package ports

import "context"

type Services struct {
	User UserServicePort
	Link LinkServicePort
}

type LinkServicePort interface {
	NewLink(ctx context.Context, link, userEmail string) (string, error)
	GetDestinationByLink(ctx context.Context, link string) (string, error)
}

type UserServicePort interface {
	Signup(ctx context.Context, email, password string) error
	Signin(ctx context.Context, email, password string) (string, error)
}
