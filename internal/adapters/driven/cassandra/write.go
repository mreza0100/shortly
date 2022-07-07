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
	const cql = `INSERT INTO users (email, password) VALUES (?, ?)`
	return w.session.Query(cql, user.Email, user.Password).Exec()
}

func (w *cassandraWrite) SaveLink(_ context.Context, short, destination, email string) error {
	const cql = `
		INSERT INTO links
		(short, destination, user_email, created_at)
		VALUES (?, ?, ?, toTimestamp(now()))
	`
	return w.session.Query(cql, short, destination, email).Exec()
}

func (w *cassandraWrite) deleteCounter(_ context.Context) error {
	const cql = `TRUNCATE counter`
	return w.session.Query(cql).Exec()
}

func (w *cassandraWrite) insertCounter(_ context.Context, newCounter int64) error {
	const cql = `INSERT INTO counter (counter) VALUES (?)`
	return w.session.Query(cql, newCounter).Exec()
}

func (w *cassandraWrite) UpdateCounter(ctx context.Context, newCounter int64) error {
	// TODO: transaction
	if err := w.deleteCounter(ctx); err != nil {
		return err
	}
	return w.insertCounter(ctx, newCounter)
}
