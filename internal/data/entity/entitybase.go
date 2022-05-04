package entity

import (
	"time"
)

type EntityBase struct {
	CreatedBy   *string
	CreatedDate *time.Time
	UpdatedBy   *string
	UpdatedDate *time.Time
}
