package application

import "errors"

var (
	// ErrInvalidDType invalid datatype provided
	ErrInvalidDType = errors.New("invalid data type provided for assessment")
)
