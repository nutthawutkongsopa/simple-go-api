package model

import "github.com/google/uuid"

type CustomerSearchRequest struct {
	Name *string `form:"name"`
	CriteriaBase
}

type CustomerDataReponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type CustomerUpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
