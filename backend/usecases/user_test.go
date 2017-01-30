package usecases_test

import (
	"github.com/joaodias/time-tracker/backend/usecases"
	"github.com/joaodias/time-tracker/backend/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userName = "Joao Dias"
var email = "diasjoaoac@gmail.com"
var accessToken = "someToken"

func TestNewUserSuccess(t *testing.T) {
	mockUserRepository := &mocks.UserRepository{
		IsError: false,
	}
	userInteractor := &usecases.UserInteractor{
		UserRepository: mockUserRepository,
	}
	newUser, err := userInteractor.New(userName, email, accessToken)
	assert.Nil(t, err, "Error should be nil when the user is successfully created.")
	assert.NotNil(t, newUser)
	assert.False(t, mockUserRepository.IsError)
	assert.True(t, mockUserRepository.NewCalled)
	assert.NotEmpty(t, newUser.ID)
	assert.NotEmpty(t, newUser.AccessToken)
	assert.Equal(t, userName, newUser.Name)
	assert.Equal(t, email, newUser.Email)
}
