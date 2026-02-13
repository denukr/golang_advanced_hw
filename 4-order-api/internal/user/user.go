package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone     string         `json:"phone" gorm:"uniqueIndex:idx_phone_deleted"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:idx_phone_deleted"`
	SessionId string         `json:"sessionId" gorm:"uniqueIndex"`
}

func NewUser(phone, sessionId string) *User {
	return &User{
		Phone:     phone,
		SessionId: sessionId,
	}
}
