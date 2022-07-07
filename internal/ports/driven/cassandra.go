package driven

import (
	"context"

	"github.com/mreza0100/shortly/internal/models"
)

type CassandraReadPort interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetDestinationByLink(_ context.Context, key string, isCustom bool) (string, error)
	GetCounter(_ context.Context) (int64, error)
}

type CassandraWritePort interface {
	UserSignup(ctx context.Context, user *models.User) error
	SaveLink(ctx context.Context, short, destination, email string) error
	UpdateCounter(_ context.Context, newCounter int64) error
}
