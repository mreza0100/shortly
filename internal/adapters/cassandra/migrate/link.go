package migrate

import "github.com/gocql/gocql"

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
