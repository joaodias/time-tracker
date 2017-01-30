package usecases

import (
	domain "github.com/joaodias/time-tracker/backend/domain"
	"github.com/pborman/uuid"
)

// TimeSession represents a time tracker session. At first glance this might
// seem a bit redundant when looking at the domain timesession entity. However
// it makes sense since the domain entities should just be afected by the
// interactor and not by the outter layers.
type TimeSession struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Duration         int    `json:"duration"`
	UserID           string `json:"userId"`
	WantCalendar     bool   `json:"wantCalendar"`
	InitialTimestamp string `json:"initialTimestamp"`
	CreatedAt        string `json:"createdAt"`
}

// TimeSessionInteractor handles the time session use cases providing it with
// necessary dependencies.
type TimeSessionInteractor struct {
	TimeSessionRepository domain.TimeSessionRepository
	Logger                Logger
}

// Logger abstracts the logging of messages.
type Logger interface {
	Log(message string) error
}

// New creates a new time tracker session in the repository.
func (interactor *TimeSessionInteractor) New(name string, duration int, userID string, wantCalendar bool, initialTimestamp string) (*TimeSession, error) {
	domainTimeSession := domain.TimeSession{
		ID:               uuid.New(),
		Name:             name,
		Duration:         duration,
		UserID:           userID,
		WantCalendar:     wantCalendar,
		InitialTimestamp: initialTimestamp,
	}
	err := interactor.TimeSessionRepository.Store(domainTimeSession)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return nil, err
	}
	return &TimeSession{
		ID:               domainTimeSession.ID,
		Name:             domainTimeSession.Name,
		Duration:         domainTimeSession.Duration,
		UserID:           domainTimeSession.UserID,
		WantCalendar:     domainTimeSession.WantCalendar,
		InitialTimestamp: domainTimeSession.InitialTimestamp,
	}, nil
}

// List lists the time tracker sessions in the repository.
func (interactor *TimeSessionInteractor) List(period string, userID string) ([]*TimeSession, error) {
	domainTimeSessions, err := interactor.TimeSessionRepository.GetAll(period, userID)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return nil, err
	}
	timeSessions := make([]*TimeSession, len(domainTimeSessions))
	for i, domainTimeSession := range domainTimeSessions {
		timeSessions[i] = &TimeSession{
			ID:        domainTimeSession.ID,
			Name:      domainTimeSession.Name,
			Duration:  domainTimeSession.Duration,
			UserID:    domainTimeSession.UserID,
			CreatedAt: domainTimeSession.CreatedAt,
		}
	}
	return timeSessions, nil
}
