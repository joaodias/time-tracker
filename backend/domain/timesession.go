package domain

// TimeSessionRepository abstracts the storage for the time tracker session.
type TimeSessionRepository interface {
	Store(TimeSession) error
	GetAll(string, string) ([]*TimeSession, error)
}

// TimeSession represents a time tracker session.
type TimeSession struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Duration         int    `json:"duration"`
	UserID           string `json:"userId"`
	WantCalendar     bool   `json:"wantCalendar"`
	InitialTimestamp string `json:"initialTimestamp"`
	CreatedAt        string `json:"createdAt"`
}
