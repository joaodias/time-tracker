package mocks

import (
	"github.com/joaodias/time-tracker/backend/domain"
)

type UserRepository struct {
	NewCalled bool
	IsError   bool
}

func (userRepo *UserRepository) New(user domain.User) error {
	userRepo.NewCalled = true
	return nil
}
