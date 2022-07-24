package cassandrarepo

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
)

type cassandraRead struct {
	session *gocql.Session
}

func (r *cassandraRead) GetUserByEmail(_ context.Context, email string) (*models.User, error) {
	iter := r.session.Query(`SELECT * FROM shortly.user_email_index WHERE email = ? LIMIT 1`, email).Iter()
	defer iter.Close()

	m := map[string]interface{}{}
	if !iter.MapScan(m) {
		return nil, customerror.NotFound
	}

	user := &models.User{
		Email:    m["email"].(string),
		Password: m["password"].(string),
	}

	return user, nil
}

func (r *cassandraRead) GetLinkByShort(_ context.Context, short string) (*models.Link, error) {
	iter := r.session.Query(`SELECT * FROM shortly.link_short_index WHERE short = ? LIMIT 1`, short).Iter()
	defer iter.Close()

	m := map[string]interface{}{}
	if !iter.MapScan(m) {
		return nil, customerror.NotFound
	}

	link := &models.Link{
		Short:       m["short"].(string),
		Destination: m["destination"].(string),
		UserId:      m["user_id"].(string),
		CreatedAt:   m["created_at"].(time.Time),
	}

	return link, nil
}

func (r *cassandraRead) GetCounter(_ context.Context) (int64, error) {
	query := r.session.Query(`SELECT counter FROM shortly.counter LIMIT 1`)
	if err := query.Exec(); err != nil {
		return 0, err
	}
	iter := query.Iter()
	defer iter.Close()

	var counter int64
	if !iter.Scan(&counter) {
		return 0, customerror.NotFound
	}

	return counter, nil
}

func (r *cassandraRead) HealthCheck(_ context.Context) bool {
	query := r.session.Query(`SELECT now() FROM system.local`)
	if err := query.Exec(); err != nil {
		return false
	}
	iter := query.Iter()
	defer iter.Close()

	var now string
	if !iter.Scan(&now) {
		return false
	}

	return now != ""
}
