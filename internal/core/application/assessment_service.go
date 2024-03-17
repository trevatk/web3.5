package application

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/trevatk/web3.5/internal/adapter/port/database/repository/assessments"
	"github.com/trevatk/web3.5/internal/core/domain"
)

// AssessmentService
type AssessmentService struct {
	querier assessments.Querier
	broker  domain.MessageBroker
}

// NewAssessmentService
func NewAssessmentService(broker domain.MessageBroker) *AssessmentService {
	return &AssessmentService{
		broker: broker,
	}
}

// InsertValue
func (as *AssessmentService) InsertValue(ctx context.Context, value *domain.NewAssessmentValue) (*domain.AssessmentValue, error) {

	err := as.isValueOfType(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("unable to validate provided assessment value %v", err)
	}

	valuebytes, err := json.Marshal(value.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal provided value %v", err)
	}

	p := assessments.InsertAssessmentValueParams{
		PatientID:             value.PatientID,
		AssessmentID:          value.AssessmentID,
		AssessmentAttributeID: value.AssessmentAttributeID,
		Input:                 valuebytes,
	}

	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	sv, err := as.querier.InsertAssessmentValue(timeout, p)
	if err != nil {
		return nil, fmt.Errorf("failed to insert assessment value %v", err)
	}

	envbytes, err := json.Marshal(&domain.ValueEnvelope{
		AssessmentID: sv.AssessmentID.String(),
		AttributeID:  sv.AssessmentAttributeID.String(),
		ValueID:      sv.ID.String(),
		Value:        valuebytes,
		CreatedAt:    sv.CreatedAt.Time,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal value envelope %v", err)
	}

	t := domain.CreateAssessmentAttributeValue.String()
	as.broker.Publish(ctx, t, envbytes)

	return &domain.AssessmentValue{
		Value:     sv.Input,
		CreatedAt: sv.CreatedAt.Time,
	}, nil
}

// internal package function to verify provided value
// is of the assessment attribute data type
func (as *AssessmentService) isValueOfType(ctx context.Context, value *domain.NewAssessmentValue) error {

	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	p := assessments.ReadAssessmentAttributeParams{
		ID:           value.AssessmentAttributeID,
		AssessmentID: value.AssessmentID,
	}

	sa, err := as.querier.ReadAssessmentAttribute(timeout, p)
	if err != nil {
		return fmt.Errorf("failed to read assessment attribute %v", err)
	}

	v := value.Value

	if sa.Dtype == "string" {
		if _, ok := v.(string); !ok {
			return ErrInvalidDType
		}
	} else if sa.Dtype == "int32" {

	} else if sa.Dtype == "int64" {

	} else if sa.Dtype == "float32" {

	}

	return nil
}

func isType(_ any, _ string) error {
	// TODO
	// implement function
	return nil
}
