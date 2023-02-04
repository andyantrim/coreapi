package messages

import (
	"fmt"
	"net/http"
)

type ErrorType string
type ErrorCode string

// Constants for error Types
const (
	ErrorTypeClient   = "clientException"
	ErrorTypeNotFound = "clientException"
	ErrorTypeServer   = "serverException"
)

const (
	ErrorCodeRequiredField = "400-001"
	ErrorCodeValidation    = "400-002"

	ErrorCodeNotFoundDB = "404-001"

	ErrorCodeServerUnknown = "500-001"
)

type GenericError struct {
	Type     ErrorType
	Title    string
	Code     ErrorCode
	HTTPCode int
	Message  string
	OK       bool
}

func (g GenericError) Error() string {
	return fmt.Sprintf("%s: %s", g.Type, g.Title)
}

func NewRequiredFieldError(message string) *GenericError {
	return &GenericError{
		Type:     ErrorTypeClient,
		Title:    "Client execption",
		Code:     ErrorCodeRequiredField,
		HTTPCode: http.StatusBadRequest,
		Message:  message,
	}
}

func NewValidationError(message string) *GenericError {
	return &GenericError{
		Type:     ErrorTypeClient,
		Title:    "Client execption",
		Code:     ErrorCodeValidation,
		HTTPCode: http.StatusBadRequest,
		Message:  message,
	}
}

func NewNotFound(message string) *GenericError {
	return &GenericError{
		Type:     ErrorTypeClient,
		Title:    "Entity not found",
		Code:     ErrorCodeNotFoundDB,
		HTTPCode: http.StatusNotFound,
		Message:  message,
	}
}

func NewServerError(message string) *GenericError {
	return &GenericError{
		Type:     ErrorTypeServer,
		Code:     ErrorCodeServerUnknown,
		HTTPCode: http.StatusInternalServerError,
		Message:  message,
		Title:    "Unknown server error",
	}
}
