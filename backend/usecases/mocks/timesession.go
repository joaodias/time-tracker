package mocks

import (
	"github.com/joaodias/time-tracker/backend/domain"
	"github.com/pkg/errors"
)

type Logger struct{}

type TimeSessionRepository struct {
	// IsError signals the intent of throwing an error. If true there is an
	// error.
	IsError bool
	// CalledStore and CalledGetAll are used to spy on the Store and GetAll
	// methods and understand if they were called during the execution.
	CalledStore  bool
	CalledGetAll bool
}

func (repo *TimeSessionRepository) Store(domain.TimeSession) error {
	repo.CalledStore = true
	if repo.IsError {
		return errors.New("Error while storing a new time session.")
	}
	return nil
}

func (repo *TimeSessionRepository) GetAll(string) ([]*domain.TimeSession, error) {
	repo.CalledGetAll = true
	if repo.IsError {
		return nil, errors.New("Error while getting all the time sessions.")
	}
	timeSession := &domain.TimeSession{
		ID:        "Some random ID",
		Name:      "Cool time tracker session",
		Duration:  123,
		CreatedAt: "Some random date",
	}
	var timeSessions []*domain.TimeSession
	timeSessions = append(timeSessions, timeSession)
	return timeSessions, nil
}

func (logger *Logger) Log(message string) error {
	return nil
}
