// TODO: move this migration to CQL files, copy them to the Cassandra container, and run them by the Cassandra database
package migrate

import (
	"github.com/gocql/gocql"
)

const CounterID = 1

func MigrateCassandra(session *gocql.Session) error {
	var err error

	if err = migrateUsersTable(session); err != nil {
		return err
	}
	if err = migrateLinkTable(session); err != nil {
		return err
	}
	if err = migrateCounterTable(session); err != nil {
		return err
	}

	return nil
}

func migrateUsersTable(session *gocql.Session) error {
	const cql = `
	CREATE TABLE IF NOT EXISTS users
		(
			email    text,
			password text,

			PRIMARY KEY (email)
		)
	`
	return session.Query(cql).Exec()
}

func migrateLinkTable(session *gocql.Session) error {
	const cql = `
	CREATE TABLE IF NOT EXISTS links 
		(
			short text,
			destination text,

			user_email text,
			created_at timestamp,

			PRIMARY KEY ((short), user_email)
		)
	`
	return session.Query(cql).Exec()
}

func migrateCounterTable(session *gocql.Session) error {
	const createTableCmd = `
	CREATE TABLE IF NOT EXISTS counter
		(
			id         int,
			counter    int,

			PRIMARY KEY (int)
		)
	`
	if err := session.Query(createTableCmd).Exec(); err != nil {
		return err
	}

	const inserCounterCmd = `INSERT INTO counter (id, counter) VALUES (?, ?)`
	return session.Query(inserCounterCmd, CounterID, 0).Exec()
}
