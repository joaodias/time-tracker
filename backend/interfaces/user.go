package interfaces

import (
	"fmt"
	"github.com/joaodias/time-tracker/backend/domain"
)

type DatabaseUserRepository DatabaseRepository

func (repository *DatabaseUserRepository) New(user domain.User) error {
	_, err := repository.DatabaseHandler.Query(repository.buildNewUserStatement(user))
	return err
}

func (repository *DatabaseUserRepository) buildNewUserStatement(user domain.User) string {
	query := fmt.Sprintf(`INSERT INTO auth_user (id, name, email, access_token, created_at, last_login) VALUES ('%s', '%s', '%s', '%s', NOW(), NOW()) ON CONFLICT (email) DO UPDATE SET last_login = NOW(), access_token = '%s'`, user.ID, user.Name, user.Email, user.AccessToken, user.AccessToken)
	return query
}
