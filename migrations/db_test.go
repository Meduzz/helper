package migrations

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	mig *Migrator
)

func TestMain(m *testing.M) {
	dbConn, err := sql.Open("postgres", "user=test dbname=test password=test sslmode=disable")

	if err != nil {
		panic(err)
	}

	db = dbConn

	exitCode := m.Run()

	dbConn.Exec("drop table testing")
	dbConn.Exec("drop table db_migrations")
	dbConn.Exec("drop sequence migrations_seq")

	os.Exit(exitCode)
}

func Test_StartTheLib(t *testing.T) {
	migRef, err := NewMigrator(db)

	if err != nil {
		t.Fatalf("Creating migrator threw error %s", err.Error())
	}

	mig = migRef
}

func Test_TwoMigrations(t *testing.T) {
	err := mig.Run("add", "Add a table", "create table testing (note varchar(50))")
	if err == nil {
		err = mig.Run("add", "Add a note", "insert into testing (note) value ('hey hey')")
	}

	if err != nil {
		t.Errorf("Migrations threw error: %s", err.Error())
	}

	row := mig.db.QueryRow("select count(id) from db_migrations")
	count := 0
	err = row.Scan(&count)

	if err != nil {
		t.Fatalf("Fetching row count threw error %s", err.Error())
	}

	if count > 1 {
		t.Fatal("Did not expect more than 1 row in db_migrations")
	}
}

func Test_RunMoarMigrations(t *testing.T) {
	err := mig.Run("note1", "Add the first note", "insert into testing (note) values ('I was here')")
	if err == nil {
		err = mig.Run("note2", "Add the second note", "insert into testing (note) values ('Twice!')")
	}

	if err != nil {
		t.Errorf("Migrations threw error: %s", err.Error())
	}

	row := mig.db.QueryRow("select count(id) from db_migrations")
	count := 0
	err = row.Scan(&count)

	if err != nil {
		t.Fatalf("Fetching row count threw error %s", err.Error())
	}

	if count != 3 {
		t.Fatalf("Expected 3 rows in db_migrations, was: %d", count)
	}

	row = mig.db.QueryRow("select count(*) from testing")
	err = row.Scan(&count)

	if err != nil {
		t.Fatalf("Fetching row count threw error %s", err.Error())
	}

	if count != 2 {
		t.Fatalf("Expected 2 rows in testing, was: %d", count)
	}
}

func Test_BadSQL(t *testing.T) {
	err := mig.Run("note3", "Add the first note", "insert into unknown (note) values ('I was here')")

	if err == nil {
		t.Error("Expected an error")
	}

	row := mig.db.QueryRow("select count(id) from db_migrations")
	count := 0
	err = row.Scan(&count)

	if err != nil {
		t.Fatalf("Fetching row count threw error %s", err.Error())
	}

	if count != 3 {
		t.Fatalf("Expected 3 rows in db_migrations, was: %d", count)
	}
}
