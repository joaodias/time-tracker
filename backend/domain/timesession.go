package domain

// TimeSessionRepository abstracts the storage for the time tracker session.
type TimeSessionRepository interface {
	Store(TimeSession) error
	GetAll(string) ([]*TimeSession, error)
}

// TimeSession represents a time tracker session.
type TimeSession struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	CreatedAt string `json:"createdAt"`
}
