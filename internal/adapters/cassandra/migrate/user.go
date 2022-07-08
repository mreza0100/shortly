package migrate

import "github.com/gocql/gocql"

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
