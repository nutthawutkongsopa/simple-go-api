package core

type DataErrorException struct {
	ErrorCode *string
	Message   *string
	Details   *string
}

func NewDataErrorException(errorCode string, message string, details string) *DataErrorException {
	result := new(DataErrorException)
	result.Details = &details
	result.Message = &message
	result.ErrorCode = &errorCode
	return result
}
