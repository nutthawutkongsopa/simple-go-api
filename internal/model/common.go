package model

type CriteriaBase struct {
	PageLength    *int    `json:"pageLength" form:"pageLength"`
	Page          *int    `json:"page" form:"page"`
	SortBy        *string `json:"sortBy" form:"sortBy"`
	SortDirection *string `json:"sortDirection" form:"sortDirection"`
}

type SearchResult[T interface{}] struct {
	ItemCount  int `json:"itemCount"`
	ResultData []T `json:"resultData"`
}

type APIResult struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type APIErrorDetail struct {
	ErrorCode *string `json:"errorCode"`
	Message   *string `json:"message"`
	Details   *string `json:"details"`
}

func NewAPIResult(status int, success bool, data interface{}) *APIResult {
	result := new(APIResult)
	result.Data = &data
	result.Status = status
	result.Success = success
	return result
}

func NewAPIResultSuccess(data interface{}) *APIResult {
	result := new(APIResult)
	result.Data = &data
	result.Status = 200
	result.Success = true
	return result
}

func NewAPIResultError(status int, errorCode string, message string, details string) *APIResult {
	result := new(APIResult)
	detail := new(APIErrorDetail)
	detail.Details = &details
	detail.ErrorCode = &errorCode
	detail.Message = &message
	result.Data = *detail
	result.Status = status
	result.Success = false
	return result
}
