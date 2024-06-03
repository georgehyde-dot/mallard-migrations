package db_connection

import (
	"context"
	"database/sql"
	"fmt"
)

type DBType string

const (
	Postgres DBType = "postgres"
	MSSQL    DBType = "mssql"
)

type DBConnection interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Close(ctx context.Context) error
}

// pgx is not directly compatible with the sql package, I could just use the generic connection,
// but figuring out how to switch cleanly seems more fun
// type PostgresConnection struct {
// 	conn *pgx.Conn
// }

// func (p *PostgresConnection) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
// 	return p.conn.QueryRow(ctx, query, args...)
// }

// func (p *PostgresConnection) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	return p.conn.Exec(ctx, query, args...)
// }

// func (p *PostgresConnection) Close(ctx context.Context) error {
// 	return p.conn.Close(ctx)
// }

type MSSQLConnection struct {
	conn *sql.DB
}

func (m *MSSQLConnection) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return m.conn.QueryRowContext(ctx, query, args...)
}

func (m *MSSQLConnection) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.conn.ExecContext(ctx, query, args...)
}

func (m *MSSQLConnection) Close(ctx context.Context) error {
	return m.conn.Close()
}

// connect to specified db
func ConnectToDB(dbType DBType, url string) (DBConnection, error) {
	switch dbType {
	// case Postgres:
	// 	conn, err := pgx.Connect(context.Background(), url)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &PostgresConnection{conn: conn}, nil
	case MSSQL:
		conn, err := sql.Open("sqlserver", url)
		if err != nil {
			return nil, err
		}
		return &MSSQLConnection{conn: conn}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

// postgres TODO implementation
// figure out how to pull multiple connections under the dbConnection

// func connectToPostgresDB(url string) (*pgx.Conn, error) {
// 	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())
// }
