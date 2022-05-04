package entity

import (
	"github.com/google/uuid"
)

type Product struct {
	EntityBase
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	Remark string
}
