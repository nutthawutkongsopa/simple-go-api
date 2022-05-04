package model

import "github.com/google/uuid"

type ProductSearchRequest struct {
	Name   *string `form:"name"`
	Remark *string `form:"remark"`
	CriteriaBase
}

type ProductDataReponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Remark string    `json:"remark"`
}

type ProductUpdateRequest struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}
