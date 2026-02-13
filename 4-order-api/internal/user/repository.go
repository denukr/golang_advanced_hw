package user

import (
	"errors"
	"golang-adv/4-order-api/pkg/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (repo *UserRepository) Create(phone, code, sessionId string) (*User, error) {
	var user User = User{
		Phone:     phone,
		SessionId: sessionId,
		Code:      code,
	}
	result := repo.Database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*User, error) {
	var user User
	result := repo.Database.First(&user, "phone = ?", phone)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	result := repo.Database.Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetBySessionId(sessionId string) (*User, error) {
	var user User
	result := repo.Database.First(&user, "session_id = ?", sessionId)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
