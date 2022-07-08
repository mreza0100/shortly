package cassandrarepo

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/mreza0100/shortly/internal/models"
)

type cassandraWrite struct {
	session *gocql.Session
}

func (w *cassandraWrite) UserSignup(_ context.Context, user *models.User) error {
	return w.session.Query(`INSERT INTO users (email, password) VALUES (?, ?)`, user.Email, user.Password).Exec()
}

func (w *cassandraWrite) SaveLink(_ context.Context, short, destination, email string) error {
	return w.session.Query(`
		INSERT INTO links
		(short, destination, user_email, created_at)
		VALUES (?, ?, ?, toTimestamp(now()))
	`, short, destination, email).Exec()
}

func (w *cassandraWrite) deleteCounter(_ context.Context) error {
	// In cassandra, there is no way to delete a row by primary key or without a where clause.
	return w.session.Query(`TRUNCATE counter`).Exec()
}

func (w *cassandraWrite) insertCounter(_ context.Context, newCounter int64) error {
	return w.session.Query(`INSERT INTO counter (counter) VALUES (?)`, newCounter).Exec()
}

func (w *cassandraWrite) UpdateCounter(ctx context.Context, newCounter int64) error {
	// TODO: a transaction would be more efficient. error is rare here, but possible.
	if err := w.deleteCounter(ctx); err != nil {
		return err
	}
	return w.insertCounter(ctx, newCounter)
}
