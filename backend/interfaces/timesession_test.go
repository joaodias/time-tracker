package interfaces_test

import (
	"github.com/joaodias/time-tracker/backend/domain"
	"github.com/joaodias/time-tracker/backend/interfaces"
	"github.com/joaodias/time-tracker/backend/interfaces/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ExpectedStoreStatement       = "INSERT INTO time_session (id, name, duration, created_at) VALUES ('Some random ID', 'Cool time tracker session', '123', NOW())"
	ExpectedGetAllStatementDay   = "SELECT * FROM time_session WHERE created_at > NOW() - INTERVAL '1 days'"
	ExpectedGetAllStatementWeek  = "SELECT * FROM time_session WHERE created_at > NOW() - INTERVAL '1 weeks'"
	ExpectedGetAllStatementMonth = "SELECT * FROM time_session WHERE created_at > NOW() - INTERVAL '1 months'"
)

func TestStoreStorageFail(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{
		IsError: true,
	}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSession := domain.TimeSession{}
	err := mockDatabaseTimeSessionRepository.Store(timeSession)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.NotNil(t, err)
}

func TestStoreStorageSuccess(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSession := domain.TimeSession{
		ID:       mocks.IDContent,
		Name:     mocks.NameContent,
		Duration: mocks.DurationContent,
	}
	err := mockDatabaseTimeSessionRepository.Store(timeSession)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.Nil(t, err)
	assert.Equal(t, ExpectedStoreStatement, mockDatabaseHandler.ExecutedStatement)
}

func TestGetAllStorageFail(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{
		IsError: true,
	}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSessions, err := mockDatabaseTimeSessionRepository.GetAll(interfaces.Day)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.Nil(t, timeSessions)
	assert.NotNil(t, err)
}

func TestGetAllRowScanFailed(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{
		IsRowsScanError: true,
	}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSessions, err := mockDatabaseTimeSessionRepository.GetAll(interfaces.Day)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.NotNil(t, err)
	assert.Equal(t, mocks.ScanError, err.Error())
	assert.Nil(t, timeSessions)
}

func TestGetAllSuccessDay(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSessions, err := mockDatabaseTimeSessionRepository.GetAll(interfaces.Day)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.Equal(t, ExpectedGetAllStatementDay, mockDatabaseHandler.ExecutedStatement)
	assert.Nil(t, err)
	assert.NotNil(t, timeSessions)
}

func TestGetAllSuccessWeek(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSessions, err := mockDatabaseTimeSessionRepository.GetAll(interfaces.Week)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.Equal(t, ExpectedGetAllStatementWeek, mockDatabaseHandler.ExecutedStatement)
	assert.Nil(t, err)
	assert.NotNil(t, timeSessions)
}

func TestGetAllSuccessMonth(t *testing.T) {
	mockDatabaseHandler := &mocks.DatabaseHandler{}
	mockDatabaseTimeSessionRepository := &interfaces.DatabaseTimeSessionRepository{}
	mockDatabaseTimeSessionRepository.DatabaseHandler = mockDatabaseHandler
	timeSessions, err := mockDatabaseTimeSessionRepository.GetAll(interfaces.Month)
	assert.True(t, mockDatabaseHandler.QueryCalled)
	assert.Equal(t, ExpectedGetAllStatementMonth, mockDatabaseHandler.ExecutedStatement)
	assert.Nil(t, err)
	assert.NotNil(t, timeSessions)
}
