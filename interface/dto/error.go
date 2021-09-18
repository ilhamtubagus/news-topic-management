package dto

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}
func NewBadRequestError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusBadRequest,
	}
}
