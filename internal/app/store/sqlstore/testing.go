package sqlstore

import (
	"strings"
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func (... string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRANCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}