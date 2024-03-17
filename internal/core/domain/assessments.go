package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// NewAssessmentValue
type NewAssessmentValue struct {
	PatientID             uuid.UUID
	AssessmentID          uuid.UUID
	AssessmentAttributeID uuid.UUID
	Value                 any
}

// AssessmentValue
type AssessmentValue struct {
	ID                    uuid.UUID
	PatientID             uuid.UUID
	AssessmentID          uuid.UUID
	AssessmentAttributeID uuid.UUID
	Value                 any
	CreatedAt             time.Time
}

// Assessments
type Assessments interface {
	// InsertValue
	InsertValue(ctx context.Context, newValue *NewAssessmentValue) (*AssessmentValue, error)
}
