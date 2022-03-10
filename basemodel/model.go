package basemodel

import (
	"fmt"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt *LocalTime `json:"created_at"`
	UpdatedAt *LocalTime `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
