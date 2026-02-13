package auth

import (
	"errors"
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
	if errors.Is(err, userPkg.ErrRecordNotFound) {
		user, err = service.UserRepository.Create(phone, GenerateCode(4), sessionId)
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
		Code:      GenerateCode(4),
	})
	if err != nil {
		return "", err
	}
	return user.SessionId, nil
}
