package cassandrarepo

import (
	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/entities"
)

type cassandraRead struct {
	session *gocql.Session
}

func (r cassandraRead) GetUserByEmail(email string) (*entities.User, error) {
	const cql = `SELECT * FROM users WHERE email = ? LIMIT 1`
	iter := r.session.Query(cql, email).Iter()

	users := make([]entities.User, 0, 1)

	m := map[string]interface{}{}
	for iter.MapScan(m) {
		users = append(users, entities.User{
			Email:    m["email"].(string),
			Password: m["password"].(string),
		})
		m = map[string]interface{}{}
	}

	if len(users) <= 0 {
		return nil, nil
	}
	return &users[0], nil
}
