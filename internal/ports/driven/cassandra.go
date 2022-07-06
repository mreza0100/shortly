package driven

import "github.com/mreza0100/shortly/internal/models"

type CassandraReadPort interface {
	GetUserByEmail(email string) (*models.User, error)
}

type CassandraWritePort interface {
	UserSignup(user *models.User) error
}
