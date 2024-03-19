package domain

import (
	"context"

	"github.com/google/uuid"
)

// NewSurvey
type NewSurvey struct {
	Name        string
	Description string
}

// Survey
type Survey struct {
	ID          uuid.UUID
	Name        string
	Description string
}

// NewAttribute
type NewAttribute struct{}

// Attribute
type Attribute struct{}

// NewValue
type NewValue struct{}

// Value
type Value struct{}

// Surveys
type Surveys interface {
	// Create
	Create(ctx context.Context, newSurvey *NewSurvey) (*Survey, error)
	// CreateAttribute
	CreateAttribute(ctx context.Context, newAttr *NewAttribute) (*Attribute, error)
	// CreateValue
	CreateValue(ctx context.Context, newValue *NewValue) (*Value, error)
}
