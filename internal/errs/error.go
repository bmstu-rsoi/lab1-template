package errs

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type ValidationErrorResponse struct {
	Message string `json:"message"`
	Errors  struct {
		AdditionalProperties string `json:"additionalProperties"`
	} `json:"errors"`
}
