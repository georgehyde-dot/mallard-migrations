package trackmigration

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// type migration struct {
// 	msg            string
// 	status         bool
// 	execution_time string
// 	duration       time.Duration
// }

// type migrationDBRow struct {
// 	id int
// 	migration
// }

func InitializeMigrationsDB(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error initializing migrations db %v", err)
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS mallard_migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			query TEXT NOT NULL,
			status BOOL NOT NULL,
			execution_time DATETIME NOT NULL,
			duration INTEGER
		)`,
	)
	if err != nil {
		return fmt.Errorf("error creating migrations table %v", err)
	}
	return nil
}

func TrackMigation(r *sql.Result, msg, dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error connecting to migrations db %v", err)
	}
	results, err := db.Exec(`INSERT INTO mallard_migrations (query, result, status, execution_time, duration) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("error inserting migration %v", err)
	}
	fmt.Printf("Results: %v\n", results)
	return nil
}
