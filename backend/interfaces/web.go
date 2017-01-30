package interfaces

import (
	"encoding/json"
	"github.com/joaodias/time-tracker/backend/usecases"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"strconv"
)

// TimeSessionInteractor is the interactor that handles the domain manipulation
// for the time session.
type TimeSessionInteractor interface {
	New(string, int, string, bool, string) (*usecases.TimeSession, error)
	List(string, string) ([]*usecases.TimeSession, error)
}

type UserInteractor interface {
	New(string, string, string) (*usecases.User, error)
}

// WebHandler handles the request provinient from the web.
type WebHandler struct {
	TimeSessionInteractor TimeSessionInteractor
	UserInteractor        UserInteractor
	Logger                usecases.Logger
}

// NewTimeSession creates a new time tracker session.
func (web *WebHandler) NewTimeSession(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	userID := r.FormValue("userId")
	initialTimestamp := r.FormValue("initialTimestamp")
	wantCalendar, err := strconv.ParseBool(r.FormValue("gCalendar"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	web.Logger.Log("userId: " + userID)
	timeSession, err := web.TimeSessionInteractor.New(name, duration, userID, wantCalendar, initialTimestamp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshalledTimeSession, err := json.Marshal(timeSession)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	web.Logger.Log("New Time Session: " + string(marshalledTimeSession))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledTimeSession)
}

// ListTimeSessions lists the time tracker sessions.
func (web *WebHandler) ListTimeSessions(w http.ResponseWriter, r *http.Request) {
	period := r.FormValue("period")
	userID := r.FormValue("userId")
	if !isPeriodValid(period) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timeSessions, err := web.TimeSessionInteractor.List(period, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshalledTimeSessions, err := json.Marshal(timeSessions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	web.Logger.Log("List time sessions: " + string(marshalledTimeSessions))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledTimeSessions)
}

func (web *WebHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	accessToken := r.FormValue("accessToken")
	user, err := web.UserInteractor.New(name, email, accessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshalledUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	web.Logger.Log("New user: " + string(marshalledUser))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledUser)
}

func (web *WebHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := authenticate(code)
	if err != nil {
		web.Logger.Log(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	type ClientToken struct {
		AccessToken string
	}
	clientToken := &ClientToken{
		AccessToken: token.AccessToken,
	}
	marshaledClientToken, err := json.Marshal(clientToken)
	if err != nil {
		web.Logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	web.Logger.Log("Authenticated Client with " + string(marshaledClientToken))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledClientToken)
}

func authenticate(code string) (*oauth2.Token, error) {
	ctx := context.Background()
	scope := make([]string, 1)
	scope[0] = "https://www.googleapis.com/auth/calendar"
	config := &oauth2.Config{
		ClientID:     "612927008159-ftu0ijkhk41a8coiil2psvcksei1r49h.apps.googleusercontent.com",
		ClientSecret: "VQdVHJZVsDJsy-Vq3Um3o2_-",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://thetimetracker.surge.sh",
		Scopes:       scope,
	}
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, err
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
