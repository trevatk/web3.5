package controller

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/trevatk/web3.5/internal/core/domain"
)

// Assessments
type Assessments struct {
	service domain.Assessments
}

// NewAssessmentValuePayload
type NewAssessmentValuePayload struct {
	PatientID string `json:"patient_id"`
}

// CreateValueParams
type CreateValueParams struct {
	Payload *NewAssessmentValuePayload `json:"payload"`
}

// AssessmentValuePayload
type AssessmentValuePayload struct {
	ID string `json:"id"`
}

// CreateValueResponse
type CreateValueResponse struct {
	Payload *AssessmentValuePayload `json:"payload"`
	Elapsed int64                   `json:"elapsed"`
}

// CreateValue
func (as *Assessments) CreateValue(c echo.Context) error {

	start := time.Now()

	var params CreateValueParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request parameters")
	}

	patientID, err := uuid.Parse(params.Payload.PatientID)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request parameters")
	}

	nv := &domain.NewAssessmentValue{
		PatientID: patientID,
	}

	ctx := c.Request().Context()

	v, err := as.service.InsertValue(ctx, nv)
	if err != nil {
		return c.String(http.StatusInternalServerError, "unable to create new assessment value")
	}

	r := &CreateValueResponse{
		Payload: &AssessmentValuePayload{
			ID: v.ID.String(),
		},
		Elapsed: time.Since(start).Milliseconds(),
	}

	return c.JSON(http.StatusCreated, r)
}
