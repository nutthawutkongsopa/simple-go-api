package model

import "github.com/google/uuid"

type CustomerSearchRequest struct {
	Name *string
	CriteriaBase
}

type CustomerDataReponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}
