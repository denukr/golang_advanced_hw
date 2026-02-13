package session

import (
	"crypto/rand"
	"encoding/base64"
)

type SessionGenerator struct {
	length int
}

func NewSessionGenerator(length int) *SessionGenerator {
	return &SessionGenerator{length: length}
}

func (s *SessionGenerator) GenerateSessinID() (string, error) {
	b := make([]byte, s.length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
