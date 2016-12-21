package infrastructure

import (
	"database/sql"
	"github.com/joaodias/time-tracker/backend/interfaces"
	// lib/pq is imported with _ because it actually runs before the main()
	// function. It regists a driver that the application will then use.
	_ "github.com/lib/pq"
)

const (
	DriverName = "postgres"
)

// PostgresHandler is the handler for the Postgres database operations.
type PostgresHandler struct {
	Conn *sql.DB
}

// PostgresRows are the table rows.
type PostgresRows struct {
	Rows *sql.Rows
}

// Query queries the database and get the needed rows.
func (handler *PostgresHandler) Query(statement string) (interfaces.Rows, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		return &PostgresRows{}, err
	}
	postgresRows := &PostgresRows{}
	postgresRows.Rows = rows
	return postgresRows, nil
}

// Scan scans the row.
func (r *PostgresRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next advance to the next row.
func (r *PostgresRows) Next() bool {
	return r.Rows.Next()
}

// Close closes the rows.
func (r *PostgresRows) Close() error {
	return r.Rows.Close()
}

// Err returns an error if any. Err may be called after an explicit or implicit
// Close.
func (r *PostgresRows) Err() error {
	return r.Rows.Err()
}

// NewPostgresHandler creates a new postgres handler.
func NewPostgresHandler(dataSource string) *PostgresHandler {
	conn, _ := sql.Open(DriverName, dataSource)
	postgresHandler := &PostgresHandler{}
	postgresHandler.Conn = conn
	return postgresHandler
}
