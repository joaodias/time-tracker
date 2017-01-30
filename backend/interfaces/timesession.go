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

type CalendarHandler interface {
	CreateEvent(string, int, string, string) error
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
	CalendarHandler CalendarHandler
}

// DatabaseTimeSessionRepository is the implementation of a database repository
// for the time session entity.
type DatabaseTimeSessionRepository DatabaseRepository

// Store is the implementation for the database repository Store method related
// to the TimeSession entity.
func (repository *DatabaseTimeSessionRepository) Store(timeSession domain.TimeSession) error {
	_, err := repository.DatabaseHandler.Query(repository.buildStoreStatement(timeSession))
	if err != nil {
		return err
	}
	// TODO: Make this the right way
	if timeSession.WantCalendar {
		accessToken, err := repository.getUserAccessToken(timeSession.UserID)
		if err != nil {
			return err
		}
		err = repository.CalendarHandler.CreateEvent(timeSession.Name, timeSession.Duration, timeSession.InitialTimestamp, accessToken)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll is the implementation for the database repository GetAll method
// related to the TimeSession entity.
func (repository *DatabaseTimeSessionRepository) GetAll(period string, userID string) ([]*domain.TimeSession, error) {
	rows, err := repository.DatabaseHandler.Query(repository.buildGetAllStatement(period, userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeSessions := make([]*domain.TimeSession, 0)
	for rows.Next() {
		timeSession := &domain.TimeSession{}
		err := rows.Scan(&timeSession.ID, &timeSession.Name, &timeSession.Duration, &timeSession.CreatedAt, &timeSession.UserID, &timeSession.WantCalendar, &timeSession.InitialTimestamp)
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

// TODO: This should no tbe her. Move to user.
func (repository *DatabaseTimeSessionRepository) getUserAccessToken(userID string) (string, error) {
	rows, err := repository.DatabaseHandler.Query(repository.buildGetAccessTokenStatement(userID))
	if err != nil {
		return "", err
	}
	defer rows.Close()
	rows.Next()
	var accessToken string
	err = rows.Scan(&accessToken)
	if err != nil {
		return "", err
	}
	if err = rows.Err(); err != nil {
		return "", err
	}
	return accessToken, nil
}

func (repository *DatabaseTimeSessionRepository) buildGetAccessTokenStatement(userID string) string {
	query := fmt.Sprintf(`SELECT access_token FROM auth_user WHERE email='%s'`, userID)
	return query
}

func (repository *DatabaseTimeSessionRepository) buildStoreStatement(timeSession domain.TimeSession) string {
	query := fmt.Sprintf(`INSERT INTO time_session (id, name, duration, user_id, want_calendar, initial_timestamp, created_at) VALUES ('%s', '%s', '%d', '%s', '%t','%s', NOW())`, timeSession.ID, timeSession.Name, timeSession.Duration, timeSession.UserID, timeSession.WantCalendar, timeSession.InitialTimestamp)
	return query
}

func (repository *DatabaseTimeSessionRepository) buildGetAllStatement(period string, userID string) string {
	var queryTimeOption string
	switch period {
	case Day:
		queryTimeOption = DayOption
	case Week:
		queryTimeOption = WeekOption
	case Month:
		queryTimeOption = MonthOption
	}
	query := fmt.Sprintf(`SELECT * FROM time_session WHERE user_id='%s' AND created_at > NOW() - INTERVAL '1 %s'`, userID, queryTimeOption)
	return query
}
