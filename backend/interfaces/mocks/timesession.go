package mocks

import (
	"errors"
	"github.com/joaodias/time-tracker/backend/interfaces"
)

const (
	ID              = 0
	Name            = 1
	Duration        = 2
	CreatedAt       = 4
	NumberOfRows    = 2
	QueryError      = "Error querying the database"
	ScanError       = "Error while scanning database rows"
	IDContent       = "Some random ID"
	NameContent     = "Cool time tracker session"
	DurationContent = 123
)

type DatabaseHandler struct {
	IsError           bool
	IsRowsScanError   bool
	IsDbErr           bool
	IsCloseErr        bool
	ExecutedStatement string
	QueryCalled       bool
}

type Rows struct {
	IsError      bool
	numberOfRows int
}

func (db *DatabaseHandler) Query(statement string) (interfaces.Rows, error) {
	db.QueryCalled = true
	db.ExecutedStatement = statement
	if db.IsError {
		return nil, errors.New(QueryError)
	}
	return &Rows{
		IsError:      db.IsRowsScanError,
		numberOfRows: NumberOfRows,
	}, nil
}

func (rows *Rows) Scan(parameters ...interface{}) error {
	if rows.IsError {
		return errors.New(ScanError)
	}
	return nil
}

func (rows *Rows) Next() bool {
	if rows.numberOfRows > 0 {
		rows.numberOfRows--
		return true
	}
	return false
}

func (rows *Rows) Close() error {
	return nil
}

func (rows *Rows) Err() error {
	return nil
}
