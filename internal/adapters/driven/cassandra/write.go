package cassandrarepo

import (
	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/models"
)

type cassandraWrite struct {
	session *gocql.Session
}

func (w *cassandraWrite) UserSignup(user *models.User) error {
	const cql = `INSERT INTO users (email, password) VALUES (?, ?)`
	return w.session.Query(cql, user.Email, user.Password).Exec()
}
