package usecases_test

import (
	"testing"

	"github.com/joaodias/time-tracker/backend/usecases"
	"github.com/joaodias/time-tracker/backend/usecases/mocks"
	"github.com/stretchr/testify/assert"
)

// Common setup for all test cases
var mockLogger = &mocks.Logger{}
var name = "Cool time tracker session"
var duration = 123
var period = "Day"

func TestNewTimeSessionStorageFail(t *testing.T) {
	mockTimeSessionRepository := &mocks.TimeSessionRepository{
		IsError: true,
	}
	mockLogger := &mocks.Logger{}
	timeSessionInteractor := &usecases.TimeSessionInteractor{
		TimeSessionRepository: mockTimeSessionRepository,
		Logger:                mockLogger,
	}
	timeSession, err := timeSessionInteractor.New(name, duration)
	assert.True(t, mockTimeSessionRepository.CalledStore)
	assert.Nil(t, timeSession)
	assert.NotNil(t, err)
}

func TestNewTimeSessionSuccess(t *testing.T) {
	mockTimeSessionRepository := &mocks.TimeSessionRepository{}
	mockLogger := &mocks.Logger{}
	timeSessionInteractor := &usecases.TimeSessionInteractor{
		TimeSessionRepository: mockTimeSessionRepository,
		Logger:                mockLogger,
	}
	timeSession, err := timeSessionInteractor.New(name, duration)
	assert.Nil(t, err, "Error should be nil when the time session is successfully created.")
	assert.True(t, mockTimeSessionRepository.CalledStore)
	assert.NotEmpty(t, timeSession.ID)
	assert.Equal(t, "Cool time tracker session", timeSession.Name)
	assert.Equal(t, 123, timeSession.Duration)
}

func TestListTimeSessionsStorageFail(t *testing.T) {
	mockTimeSessionRepository := &mocks.TimeSessionRepository{
		IsError: true,
	}
	mockLogger := &mocks.Logger{}
	timeSessionInteractor := &usecases.TimeSessionInteractor{
		TimeSessionRepository: mockTimeSessionRepository,
		Logger:                mockLogger,
	}
	timesessions, err := timeSessionInteractor.List(period)
	assert.True(t, mockTimeSessionRepository.CalledGetAll)
	assert.Nil(t, timesessions)
	assert.NotNil(t, err)
}

func TestListTimeSessionSuccess(t *testing.T) {
	mockTimeSessionRepository := &mocks.TimeSessionRepository{}
	mockLogger := &mocks.Logger{}
	timeSessionInteractor := &usecases.TimeSessionInteractor{
		TimeSessionRepository: mockTimeSessionRepository,
		Logger:                mockLogger,
	}
	timeSessions, err := timeSessionInteractor.List(period)
	assert.True(t, mockTimeSessionRepository.CalledGetAll)
	assert.NotNil(t, timeSessions)
	assert.Nil(t, err)
}
