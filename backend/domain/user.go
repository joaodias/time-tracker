package domain

// UserRepository abstrats the storage for the user.
type UserRepository interface {
	New(User) error
}

// User is the user of the system.
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"email"`
	CreatedAt   string `json:"CreatedAt"`
}
