package user

import "gorm.io/gorm"

type RecordNotFound string

func (e *RecordNotFound) Error() string {
	return gorm.ErrRecordNotFound.Error()
}

var ErrRecordNotFound *RecordNotFound
