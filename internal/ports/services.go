package ports

import "context"

type Services struct {
	User   UserServicePort
	Link   LinkServicePort
	Health HealthServicePort
}

type HealthServicePort interface {
	CheckHealth(ctx context.Context) bool
}

type LinkServicePort interface {
	NewLink(ctx context.Context, destination string, userId string) (string, error)
	GetDestinationByLink(ctx context.Context, link string) (string, error)
}

type UserServicePort interface {
	Signup(ctx context.Context, email, password string) error
	Signin(ctx context.Context, email, password string) (string, error)
}
