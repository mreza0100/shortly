package driven

import "github.com/mreza0100/shortly/internal/entities"

type CassandraReadPort interface {
	GetUserByEmail(email string) (*entities.User, error)
}

type CassandraWritePort interface {
	UserSignup(user *entities.User) error
}
