package migrate

import (
	"github.com/gocql/gocql"
)

func insertCounterRow(session *gocql.Session) error {
	const inserCounterCmd = `INSERT INTO counter (counter) VALUES (?)`
	return session.Query(inserCounterCmd, 1).Exec()
}

func isCounterTableEmpty(session *gocql.Session) (bool, error) {
	const existsCQL = `SELECT EXISTS (SELECT 1 FROM counter LIMIT 1)`
	iter := session.Query(existsCQL).Iter()

	var exists bool
	iter.Scan(&exists)

	return !exists, nil
}

func migrateCounterTable(session *gocql.Session) error {
	const createTableCmd = `
	CREATE TABLE IF NOT EXISTS counter
		(
			counter BIGINT,

			PRIMARY KEY (counter)
		)
	`
	if err := session.Query(createTableCmd).Exec(); err != nil {
		return err
	}

	isEmpty, err := isCounterTableEmpty(session)
	if err != nil {
		return err
	}
	if isEmpty {
		if err := insertCounterRow(session); err != nil {
			return err
		}
	}

	return nil
}
