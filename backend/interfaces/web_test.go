package interfaces_test

import (
	"github.com/joaodias/time-tracker/backend/interfaces"
	"github.com/joaodias/time-tracker/backend/interfaces/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewTimeSessionBadRequest(t *testing.T) {
	// The type of the request doesn't really matter here. The type of the
	// request is handled by the router. However for consistency the request
	// types are represented as they are in reality.
	request, err := http.NewRequest("POST", "/timesession?name=Cool time tracker&duration=should be int", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.NewTimeSession(recorder, request)
	assert.False(t, mockTimeSessionInteractor.NewCalled)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestNewTimeSessionInternalError(t *testing.T) {
	request, err := http.NewRequest("POST", "/timesession?name=Cool time tracker&duration=123&date=11/12/2016", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{
		IsNewError: true,
	}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.NewTimeSession(recorder, request)
	assert.True(t, mockTimeSessionInteractor.NewCalled)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestNewTimeSessionOK(t *testing.T) {
	request, err := http.NewRequest("POST", "/timesession?name=Cool time tracker&duration=123&date=11/12/2016", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.NewTimeSession(recorder, request)
	assert.True(t, mockTimeSessionInteractor.NewCalled)
	assert.Equal(t, http.StatusOK, http.StatusOK)
}

func TestListTimeSessionsBadRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/timesession?period=SomethingWrong", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{
		IsListError: true,
	}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.ListTimeSessions(recorder, request)
	assert.False(t, mockTimeSessionInteractor.ListCalled)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestListTimeSessionsInternalError(t *testing.T) {
	request, err := http.NewRequest("GET", "/timesession?period=Day", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{
		IsListError: true,
	}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.ListTimeSessions(recorder, request)
	assert.True(t, mockTimeSessionInteractor.ListCalled)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func TestListTimeSessionOK(t *testing.T) {
	request, err := http.NewRequest("GET", "/timesession?period=Day", nil)
	if err != nil {
		t.Error("Error creating request.")
	}
	recorder := httptest.NewRecorder()

	mockTimeSessionInteractor := &mocks.TimeSessionInteractor{}
	webHandler := &interfaces.WebHandler{
		TimeSessionInteractor: mockTimeSessionInteractor,
	}
	webHandler.ListTimeSessions(recorder, request)
	assert.True(t, mockTimeSessionInteractor.ListCalled)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
