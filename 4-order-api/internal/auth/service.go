package auth

import (
	"errors"
	"fmt"
	userPkg "golang-adv/4-order-api/internal/user"
	"golang-adv/4-order-api/pkg/session"

	"gorm.io/gorm"
)

type AuthService struct {
	*userPkg.UserRepository
}

func NewAuthService(userRepo *userPkg.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepo,
	}
}

func (service *AuthService) AuthByPhone(phone string) (string, error) {
	sessionGenerator := session.NewSessionGenerator(16)
	sessionId, err := sessionGenerator.GenerateSessinID()
	if err != nil {
		return "", err
	}
	user, err := service.UserRepository.GetByPhone(phone)
	fmt.Println(errors.Is(err, userPkg.ErrRecordNotFound))
	if errors.Is(err, userPkg.ErrRecordNotFound) {
		fmt.Println("Here")
		user, err = service.UserRepository.Create(phone, sessionId)
		if err != nil {
			return "", err
		}
		return user.SessionId, nil
	}
	if err != nil {
		return "", err
	}
	user, err = service.UserRepository.Update(&userPkg.User{
		Model:     gorm.Model{ID: user.ID},
		SessionId: sessionId,
	})
	if err != nil {
		return "", err
	}
	return user.SessionId, nil
}
