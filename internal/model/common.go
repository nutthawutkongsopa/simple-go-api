package model

type CriteriaBase struct {
	PageLength    *int
	Page          *int
	SortBy        *string
	SortDirection *string
}

type SearchResult[T interface{}] struct {
	ItemCount  int64
	ResultData []T
}
