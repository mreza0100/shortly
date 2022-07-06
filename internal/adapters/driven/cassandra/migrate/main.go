// TODO: move this migration to CQL files, copy them to the Cassandra container, and run them by the Cassandra database
package migrate

import (
	"github.com/gocql/gocql"
)

func MigrateCassandra(session *gocql.Session) error {
	var err error

	if err = migrateUsersTable(session); err != nil {
		return err
	}
	if err = migrateLinkTable(session); err != nil {
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

			username text,

			created_at timestamp,
			expires_at timestamp,

			PRIMARY KEY ((short), username, expires_at)
		)
	`
	return session.Query(cql).Exec()
}
