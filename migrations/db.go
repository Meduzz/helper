package migrations

import "database/sql"

type Migrator struct {
	db *sql.DB
}

const (
	sequence = "create sequence if not exists migrations_seq start 1"
	table    = "create table if not exists db_migrations (id bigint primary key default nextval('migrations_seq'), uid varchar(100), description varchar(512))"
)

// NewMigrator - creates a migrator and runs the libs owns migrations.
func NewMigrator(db *sql.DB) (*Migrator, error) {
	_, err := db.Exec(sequence)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(table)

	if err != nil {
		return nil, err
	}

	return &Migrator{db}, nil
}

// Run - runs and records the migration.
func (m *Migrator) Run(uid, description, sql string) error {
	row := m.db.QueryRow("select count(id) from db_migrations where uid = $1", uid)
	count := 0
	err := row.Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		_, err = m.db.Exec(sql)

		if err != nil {
			return err
		}

		_, err = m.db.Exec("insert into db_migrations (uid, description) values ($1, $2)", uid, description)

		return err
	}

	return nil
}
