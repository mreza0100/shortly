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
	iter := r.session.Query(`SELECT * FROM users WHERE email = ? LIMIT 1`, email).Iter()

	m := map[string]interface{}{}
	if !iter.MapScan(m) {
		return nil, er.NotFound
	}

	user := &models.User{
		Email:    m["email"].(string),
		Password: m["password"].(string),
	}

	return user, nil
}

func (r *cassandraRead) GetDestinationByLink(_ context.Context, shortLink string) (string, error) {
	query := r.session.Query(`SELECT destination FROM links WHERE short = ? LIMIT 1`, shortLink)
	if err := query.Exec(); err != nil {
		return "", err
	}

	var destination string
	if !query.Iter().Scan(&destination) {
		return "", er.NotFound
	}

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
	if !query.Iter().Scan(&counter) {
		return 0, er.NotFound
	}

	return counter, nil
}
