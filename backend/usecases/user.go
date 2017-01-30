package usecases

import (
	"github.com/joaodias/time-tracker/backend/domain"
	"github.com/pborman/uuid"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"email"`
	CreatedAt   string `json:"CreatedAt"`
}

type UserInteractor struct {
	UserRepository domain.UserRepository
	Logger         Logger
}

func (userInteractor *UserInteractor) New(name string, email string, accessToken string) (*User, error) {
	domainUser := domain.User{
		ID:          uuid.New(),
		Name:        name,
		Email:       email,
		AccessToken: accessToken,
	}
	err := userInteractor.UserRepository.New(domainUser)
	if err != nil {
		userInteractor.Logger.Log(err.Error())
		return nil, err
	}
	return &User{
		ID:          domainUser.ID,
		Name:        domainUser.Name,
		Email:       domainUser.Email,
		AccessToken: domainUser.AccessToken,
	}, nil
}
