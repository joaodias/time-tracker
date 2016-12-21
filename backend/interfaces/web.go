package interfaces

import (
	"encoding/json"
	"github.com/joaodias/time-tracker/backend/usecases"
	"net/http"
	"strconv"
)

// TimeSessionInteractor is the interactor that handles the domain manipulation
// for the time session.
type TimeSessionInteractor interface {
	New(string, int) (*usecases.TimeSession, error)
	List(string) ([]*usecases.TimeSession, error)
}

// WebHandler handles the request provinient from the web.
type WebHandler struct {
	TimeSessionInteractor TimeSessionInteractor
	Logger                usecases.Logger
}

// NewTimeSession creates a new time tracker session.
func (web *WebHandler) NewTimeSession(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timeSession, err := web.TimeSessionInteractor.New(name, duration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshalledTimeSession, err := json.Marshal(timeSession)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledTimeSession)
}

// ListTimeSessions lists the time tracker sessions.
func (web *WebHandler) ListTimeSessions(w http.ResponseWriter, r *http.Request) {
	period := r.FormValue("period")
	if !isPeriodValid(period) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timeSessions, err := web.TimeSessionInteractor.List(period)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshalledTimeSessions, err := json.Marshal(timeSessions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledTimeSessions)
}

func isPeriodValid(period string) bool {
	switch period {
	case Day:
		return true
	case Week:
		return true
	case Month:
		return true
	default:
		return false
	}
}
