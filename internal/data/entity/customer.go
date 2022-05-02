package entity

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Remark    string
}
