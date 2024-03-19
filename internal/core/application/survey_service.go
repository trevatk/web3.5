package application

import (
	"context"

	"github.com/trevatk/web3.5/internal/adapter/port/database/repository/surveys"
	"github.com/trevatk/web3.5/internal/core/domain"
)

// SurveyService
type SurveyService struct {
	querier surveys.Querier
}

// interface compliance
var _ domain.Surveys = (*SurveyService)(nil)

// NewSurveyService
func NewSurveyService() *SurveyService {
	return &SurveyService{}
}

// Create
func (ss *SurveyService) Create(ctx context.Context, newSurvey *domain.NewSurvey) (*domain.Survey, error) {
	// TODO:
	// implement function
	return nil, nil
}

// CreateAttribute
func (ss *SurveyService) CreateAttribute(ctx context.Context, newAttr *domain.NewAttribute) (*domain.Attribute, error) {
	// TODO:
	// implement function
	return nil, nil
}

// CreateValue
func (ss *SurveyService) CreateValue(ctx context.Context, newValue *domain.NewValue) (*domain.Value, error) {
	// TODO:
	// implement function
	return nil, nil
}
