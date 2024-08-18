package domain

import "fmt"

// NotFoundError is a custom error type for not found entities.
type NotFoundError struct {
	Entity string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}

// NewNotFoundError creates a new NotFoundError for a given entity.
func NewNotFoundError(entity string) error {
	return &NotFoundError{Entity: entity}
}
