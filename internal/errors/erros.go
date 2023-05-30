// Package apierrors contains API Errors definitions.
package apierrors

import (
	"encoding/json"
	"fmt"
)

// ValidationError represents the errors returned during some model's validation.
type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

// NewValidationError creates a new ValidationError based on the given params.
func NewValidationError(field string, tag string) *ValidationError {
	return &ValidationError{Field: field, Tag: tag}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Tag)
}

type APIErrorOption func(err *APIError)

type APIError struct {
	source         error
	detail         string
	httpStatusCode int
}

func (e *APIError) Source() error {
	return e.source
}

func (e *APIError) Detail() string {
	return e.detail
}

func (e *APIError) HTTPStatusCode() int {
	return e.httpStatusCode
}

func (e *APIError) Error() string {
	if e.source != nil {
		return fmt.Sprintf("an error has occurred: %v", e.source.Error())
	}
	return e.detail
}

func (e *APIError) Unwrap() error {
	return e.source
}

func WithSource(source error) APIErrorOption {
	return func(err *APIError) {
		err.source = source
	}
}

func WithDetail(detail string) APIErrorOption {
	return func(err *APIError) {
		err.detail = detail
	}
}

func WithHTTPStatusCode(httpStatusCode int) APIErrorOption {
	return func(err *APIError) {
		err.httpStatusCode = httpStatusCode
	}
}

func NewAPIError(opts ...APIErrorOption) *APIError {
	err := &APIError{}
	for _, opt := range opts {
		opt(err)
	}
	return err
}

func (e *APIError) MarshalJSON() ([]byte, error) {
	err := &struct {
		Message string `json:"message"`
	}{
		Message: e.detail,
	}
	return json.Marshal(err)
}
