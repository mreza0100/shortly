package ports

import (
	"context"

	"github.com/mreza0100/shortly/internal/models"
)

type CassandraReadPort interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetDestinationByLink(ctx context.Context, key string) (string, error)
	GetCounter(ctx context.Context) (int64, error)
}

type CassandraWritePort interface {
	UserSignup(ctx context.Context, user *models.User) error
	SaveLink(ctx context.Context, short, destination, email string) error
	UpdateCounter(ctx context.Context, newCounter int64) error
}
