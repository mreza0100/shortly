package cassandrarepo

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/models"
	"github.com/mreza0100/shortly/internal/pkg/customerror"
)

type cassandraWrite struct {
	session *gocql.Session
}

func (w *cassandraWrite) UserSignup(_ context.Context, user *models.User) error {
	err := w.session.Query(`INSERT INTO users (email, password) VALUES (?, ?)`, user.Email, user.Password).Exec()
	if err != nil {
		return customerror.GeneralFailure
	}
	return nil
}

func (w *cassandraWrite) SaveLink(_ context.Context, short, destination, userId string) error {
	err := w.session.Query(`
		INSERT INTO shortly.link_short_index
		(short, destination, user_id, created_at)
		VALUES (?, ?, ?, toTimestamp(now()))
	`, short, destination, userId).Exec()
	if err != nil {
		return customerror.GeneralFailure
	}
	return nil
}

func (w *cassandraWrite) deleteCounter(_ context.Context) error {
	// In cassandra, there is no way to delete a row by primary key or without a where clause.
	return w.session.Query(`TRUNCATE shortly.counter`).Exec()
}

func (w *cassandraWrite) insertCounter(_ context.Context, newCounter int64) error {
	return w.session.Query(`INSERT INTO shortly.counter (counter) VALUES (?)`, newCounter).Exec()
}

func (w *cassandraWrite) UpdateCounter(ctx context.Context, newCounter int64) error {
	// TODO: a transaction would be more efficient. error is rare here, but possible.
	if err := w.deleteCounter(ctx); err != nil {
		return err
	}
	return w.insertCounter(ctx, newCounter)
}
