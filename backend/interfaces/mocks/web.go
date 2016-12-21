package mocks

import (
	"github.com/joaodias/time-tracker/backend/usecases"
	"github.com/pkg/errors"
)

const (
	NewTimeSessionError   = "Error creating the new time session."
	ListTimeSessionsError = "Error listing the time sessions."
)

type TimeSessionInteractor struct {
	IsNewError  bool
	IsListError bool
	NewCalled   bool
	ListCalled  bool
}

func (timeSessionInteractor *TimeSessionInteractor) New(name string, duration int) (*usecases.TimeSession, error) {
	timeSessionInteractor.NewCalled = true
	if timeSessionInteractor.IsNewError {
		return nil, errors.New(NewTimeSessionError)
	}
	return &usecases.TimeSession{}, nil
}

func (timeSessionInteractor *TimeSessionInteractor) List(periode string) ([]*usecases.TimeSession, error) {
	timeSessionInteractor.ListCalled = true
	if timeSessionInteractor.IsListError {
		return nil, errors.New(ListTimeSessionsError)
	}
	return []*usecases.TimeSession{}, nil
}
