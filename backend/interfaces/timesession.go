package interfaces

import (
	"fmt"
	"github.com/joaodias/time-tracker/backend/domain"
)

// Consts for the time period options.
const (
	Day         = "Day"
	Week        = "Week"
	Month       = "Month"
	DayOption   = "days"
	WeekOption  = "weeks"
	MonthOption = "months"
)

// DatabaseHandler is a database abstraction. The interface can get broader or
// narrower on demand.
type DatabaseHandler interface {
	Query(string) (Rows, error)
}

// Rows abstracts rows of tables in the database.
type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

// DatabaseRepository is the repository for the database.
type DatabaseRepository struct {
	DatabaseHandler DatabaseHandler
}

// DatabaseTimeSessionRepository is the implementation of a database repository
// for the time session entity.
type DatabaseTimeSessionRepository DatabaseRepository

// Store is the implementation for the database repository Store method related
// to the TimeSession entity.
func (repository *DatabaseTimeSessionRepository) Store(timeSession domain.TimeSession) error {
	_, err := repository.DatabaseHandler.Query(repository.buildStoreStatement(timeSession))
	return err
}

// GetAll is the implementation for the database repository GetAll method
// related to the TimeSession entity.
func (repository *DatabaseTimeSessionRepository) GetAll(period string) ([]*domain.TimeSession, error) {
	rows, err := repository.DatabaseHandler.Query(repository.buildGetAllStatement(period))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeSessions := make([]*domain.TimeSession, 0)
	for rows.Next() {
		timeSession := &domain.TimeSession{}
		err := rows.Scan(&timeSession.ID, &timeSession.Name, &timeSession.Duration, &timeSession.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeSessions = append(timeSessions, timeSession)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return timeSessions, nil
}

func (repository *DatabaseTimeSessionRepository) buildStoreStatement(timeSession domain.TimeSession) string {
	query := fmt.Sprintf(`INSERT INTO time_session (id, name, duration, created_at) VALUES ('%s', '%s', '%d', NOW())`, timeSession.ID, timeSession.Name, timeSession.Duration)
	return query
}

func (repository *DatabaseTimeSessionRepository) buildGetAllStatement(period string) string {
	var queryTimeOption string
	switch period {
	case Day:
		queryTimeOption = DayOption
	case Week:
		queryTimeOption = WeekOption
	case Month:
		queryTimeOption = MonthOption
	}
	query := fmt.Sprintf(`SELECT * FROM time_session WHERE created_at > NOW() - INTERVAL '1 %s'`, queryTimeOption)
	return query
}
