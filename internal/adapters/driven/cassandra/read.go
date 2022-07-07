package cassandrarepo

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/models"
	er "github.com/mreza0100/shortly/pkg/errors"
)

type cassandraRead struct {
	session *gocql.Session
}

func (r *cassandraRead) GetUserByEmail(_ context.Context, email string) (*models.User, error) {
	const cql = `SELECT * FROM users WHERE email = ? LIMIT 1`
	iter := r.session.Query(cql, email).Iter()

	users := make([]models.User, 0, 1)

	m := map[string]interface{}{}
	for iter.MapScan(m) {
		users = append(users, models.User{
			Email:    m["email"].(string),
			Password: m["password"].(string),
		})
		m = map[string]interface{}{}
	}

	if len(users) <= 0 {
		return nil, er.NotFound
	}
	return &users[0], nil
}

func (r *cassandraRead) GetDestinationByLink(_ context.Context, shortLink string) (string, error) {
	query := r.session.Query(`SELECT destination FROM links WHERE short = ? LIMIT 1`, shortLink)
	if err := query.Exec(); err != nil {
		return "", err
	}

	var destination string
	query.Iter().Scan(&destination)

	if destination == "" {
		return "", er.NotFound
	}
	return destination, nil
}

func (r *cassandraRead) GetCounter(_ context.Context) (int64, error) {
	query := r.session.Query(`SELECT counter FROM counter LIMIT 1`)
	if err := query.Exec(); err != nil {
		return 0, err
	}

	var counter int64
	query.Iter().Scan(&counter)

	return counter, nil
}
