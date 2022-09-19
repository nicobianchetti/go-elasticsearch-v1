package api

import (
	"fmt"
	"net/http"
)

const (
	MsgInvalidQueryParams = "Invalid request query params."
	MsgInvalidQURIParams  = "Invalid request uri params."
	MsgInvalidBodyData    = "Invalid request body data."
)

type CauseList []interface{}

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Cause() CauseList
	Error() string
}

type apiErr struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (c CauseList) ToString() string {
	return fmt.Sprint(c)
}

func (e apiErr) Code() string {
	return e.ErrorCode
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e apiErr) Status() int {
	return e.ErrorStatus
}

func (e apiErr) Cause() CauseList {
	return e.ErrorCause
}

func (e apiErr) Message() string {
	return e.ErrorMessage
}

func NewApiError(message string, error string, status int, cause CauseList) ApiError {
	return apiErr{message, error, status, cause}
}

func NewInvalidQueryParamsError(err error) ApiError {
	return NewError(MsgInvalidQueryParams, http.StatusBadRequest, err)
}

func NewInvalidBodyDataError(err error) ApiError {
	return NewError(MsgInvalidBodyData, http.StatusBadRequest, err)
}

func NewInvalidURLParamsError(err error) ApiError {
	return NewError(MsgInvalidQURIParams, http.StatusBadRequest, err)
}

func NewError(msg string, status int, err error) ApiError {
	causes := CauseList{err.Error()}
	return NewApiError(msg, http.StatusText(status), status, causes)
}
