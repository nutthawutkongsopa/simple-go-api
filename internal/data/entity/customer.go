package entity

import (
	"github.com/google/uuid"
)

type Customer struct {
	EntityBase
	ID        uuid.UUID `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Remark    string
}
