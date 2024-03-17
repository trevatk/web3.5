package controller

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"

	"github.com/trevatk/web3.5/internal/core/domain"
)

// Assessments http controller for assessments
type Assessments struct {
	log     *zap.SugaredLogger
	service domain.Assessments
}

// NewAssessments return new assessments
func NewAssessments(logger *zap.Logger, service domain.Assessments) *Assessments {
	return &Assessments{
		log:     logger.Sugar().Named("AssessmentController"),
		service: service,
	}
}

// RegisterRoutesV1 for assessments controller
func (as *Assessments) RegisterRoutesV1(e *echo.Group) {
	e.POST(assessmentAttributeValuePath, as.CreateValue)
}

// NewAssessmentValuePayload http new assessment value model
type NewAssessmentValuePayload struct {
	PatientID string `json:"patient_id"`
	Value     any    `json:"value"`
}

// CreateValueParams http create value model
type CreateValueParams struct {
	Payload *NewAssessmentValuePayload `json:"payload"`
}

// AssessmentValuePayload http assessment value model
type AssessmentValuePayload struct {
	ID                 string    `json:"id"`
	AssessmentID       string    `json:"assessment_id"`
	AssmentAttributeID string    `json:"assessment_attribute_id"`
	Value              any       `json:"value"`
	CreatedAt          time.Time `json:"created_at"`
}

// CreateValueResponse http create assessment value response
type CreateValueResponse struct {
	Payload *AssessmentValuePayload `json:"payload"`
	Elapsed int64                   `json:"elapsed"`
}

// CreateValue http exposed endpoint to create assessment attribute value
func (as *Assessments) CreateValue(c echo.Context) error {

	start := time.Now()

	// validate assessment id path parameter
	id1 := c.Param("assessment_id")
	assessmentID, err := uuid.Parse(id1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request path parameters")
	}

	// validate attribute id path parameter
	id2 := c.Param("attribute_id")
	attributeID, err := uuid.Parse(id2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request path parameters")
	}

	// validate request body
	var params CreateValueParams
	err = c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request parameters")
	}

	// parse string to uuid
	patientID, err := uuid.Parse(params.Payload.PatientID)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request parameters")
	}

	// transform to application logic
	nv := &domain.NewAssessmentValue{
		AssessmentID:          assessmentID,
		AssessmentAttributeID: attributeID,
		PatientID:             patientID,
		Value:                 params.Payload.Value,
	}
	ctx := c.Request().Context()

	// insert attribute value
	v, err := as.service.InsertValue(ctx, nv)
	if err != nil {
		return c.String(http.StatusInternalServerError, "unable to create new assessment value")
	}

	// create response
	r := &CreateValueResponse{
		Payload: &AssessmentValuePayload{
			ID:                 v.ID.String(),
			AssessmentID:       v.AssessmentID.String(),
			AssmentAttributeID: v.AssessmentAttributeID.String(),
			Value:              v.Value,
			CreatedAt:          v.CreatedAt,
		},
		Elapsed: time.Since(start).Milliseconds(),
	}

	return c.JSON(http.StatusCreated, r)
}
