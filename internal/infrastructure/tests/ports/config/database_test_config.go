package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func SetupInMemoryDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory DB: %v", err)
	}

	schema := `
	CREATE TABLE quotes (
		id INTEGER PRIMARY KEY,
		literary_work_id INTEGER,
		description TEXT
	);
	INSERT INTO quotes (id, literary_work_id, description) VALUES
		(1, 101, 'Quote one'),
		(2, 102, 'Quote two');
	`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("Failed to initialize schema: %v", err)
	}

	return db
}
