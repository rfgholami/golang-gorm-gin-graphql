package models

type ErrorDetail struct {
	Message       string      `json:"message"`
	StackTrace    interface{} `json:"stackTrace"`
	Code          interface{} `json:"code"`
	ExceptionCode interface{} `json:"exceptionCode"`
}
