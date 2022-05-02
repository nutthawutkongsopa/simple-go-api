package entity

import (
	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	Remark string
}
