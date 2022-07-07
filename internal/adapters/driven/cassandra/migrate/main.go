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
	if err = migrateCounterTable(session); err != nil {
		return err
	}

	return nil
}
