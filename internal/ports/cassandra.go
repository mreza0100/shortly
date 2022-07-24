package ports

import (
	"context"

	"github.com/mreza0100/shortly/internal/models"
)

// Storage Read/Write interfaces

type StorageReadPort interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetLinkByShort(ctx context.Context, short string) (*models.Link, error)
	GetCounter(ctx context.Context) (int64, error)
	HealthCheck(ctx context.Context) bool
}

type StorageWritePort interface {
	SaveUser(ctx context.Context, user *models.User) error
	SaveLink(ctx context.Context, short, destination, userId string) error
	UpdateCounter(ctx context.Context, newCounter int64) error
}
