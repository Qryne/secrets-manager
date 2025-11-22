package auth

import (
	"strings"
)

type AuthService struct {
	AuthRepo IAuthRepository
}

func (service *AuthService) InitUserSignup(email string) error {
	sanitizedEmail := strings.ToLower(strings.Trim(email, ""))
	_, err := service.AuthRepo.CreateUserByEmail(sanitizedEmail)

	if err != nil {
		return err
	}

	return nil

}
