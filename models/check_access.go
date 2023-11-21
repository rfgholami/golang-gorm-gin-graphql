package models

type CheckAccess struct {
	InputArguments interface{} `json:"inputArguments"`
	Response       interface{} `json:"response"`
	ResponseCode   int         `json:"responseCode"`
	ErrorDetail    ErrorDetail `json:"errorDetail"`
}
