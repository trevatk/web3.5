// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: surveys.sql

package surveys

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const insertSurvey = `-- name: InsertSurvey :one
INSERT INTO surveys_eav (
    display_nm,
    description
) VALUES (
    $1, $2
)
RETURNING id, display_nm, description, created_at, updated_at
`

type InsertSurveyParams struct {
	DisplayNm   string
	Description string
}

// InsertSurvey insert new survey record
func (q *Queries) InsertSurvey(ctx context.Context, arg InsertSurveyParams) (SurveysEav, error) {
	row := q.db.QueryRow(ctx, insertSurvey, arg.DisplayNm, arg.Description)
	var i SurveysEav
	err := row.Scan(
		&i.ID,
		&i.DisplayNm,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertSurveyAttribute = `-- name: InsertSurveyAttribute :one
INSERT INTO survey_attributes (
    survey_id,
    display_nm,
    description,
    dtype
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, survey_id, display_nm, description, dtype, order_execution, created_at, updated_at
`

type InsertSurveyAttributeParams struct {
	SurveyID    uuid.UUID
	DisplayNm   string
	Description pgtype.Text
	Dtype       string
}

// InsertSurveyAttribute
func (q *Queries) InsertSurveyAttribute(ctx context.Context, arg InsertSurveyAttributeParams) (SurveyAttribute, error) {
	row := q.db.QueryRow(ctx, insertSurveyAttribute,
		arg.SurveyID,
		arg.DisplayNm,
		arg.Description,
		arg.Dtype,
	)
	var i SurveyAttribute
	err := row.Scan(
		&i.ID,
		&i.SurveyID,
		&i.DisplayNm,
		&i.Description,
		&i.Dtype,
		&i.OrderExecution,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertSurveyValue = `-- name: InsertSurveyValue :one
INSERT INTO survey_values (
    submitted_by,
    survey_id,
    survey_attribute_id,
    input
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, submitted_by, survey_id, survey_attribute_id, input, created_at
`

type InsertSurveyValueParams struct {
	SubmittedBy       uuid.UUID
	SurveyID          uuid.UUID
	SurveyAttributeID uuid.UUID
	Input             []byte
}

// InsertSurveyValue insert new survey value record
func (q *Queries) InsertSurveyValue(ctx context.Context, arg InsertSurveyValueParams) (SurveyValue, error) {
	row := q.db.QueryRow(ctx, insertSurveyValue,
		arg.SubmittedBy,
		arg.SurveyID,
		arg.SurveyAttributeID,
		arg.Input,
	)
	var i SurveyValue
	err := row.Scan(
		&i.ID,
		&i.SubmittedBy,
		&i.SurveyID,
		&i.SurveyAttributeID,
		&i.Input,
		&i.CreatedAt,
	)
	return i, err
}

const listSurveyAttributes = `-- name: ListSurveyAttributes :many
SELECT
    id,
    display_nm,
    description,
    dtype,
    order_execution
FROM
    survey_attributes
WHERE
    survey_id = $1
ORDER BY order_execution ASC
`

type ListSurveyAttributesRow struct {
	ID             uuid.UUID
	DisplayNm      string
	Description    pgtype.Text
	Dtype          string
	OrderExecution int32
}

// ListSurbeyAttributes
func (q *Queries) ListSurveyAttributes(ctx context.Context, surveyID uuid.UUID) ([]ListSurveyAttributesRow, error) {
	rows, err := q.db.Query(ctx, listSurveyAttributes, surveyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSurveyAttributesRow
	for rows.Next() {
		var i ListSurveyAttributesRow
		if err := rows.Scan(
			&i.ID,
			&i.DisplayNm,
			&i.Description,
			&i.Dtype,
			&i.OrderExecution,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSurveys = `-- name: ListSurveys :many
SELECT
    id,
    display_nm
FROM
    surveys_eav
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type ListSurveysParams struct {
	Limit  int32
	Offset int32
}

type ListSurveysRow struct {
	ID        uuid.UUID
	DisplayNm string
}

// ListSurveys retrieve partial survey records
func (q *Queries) ListSurveys(ctx context.Context, arg ListSurveysParams) ([]ListSurveysRow, error) {
	rows, err := q.db.Query(ctx, listSurveys, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSurveysRow
	for rows.Next() {
		var i ListSurveysRow
		if err := rows.Scan(&i.ID, &i.DisplayNm); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readSurvey = `-- name: ReadSurvey :one
SELECT  
    id, display_nm, description, created_at, updated_at
FROM
    surveys_eav
WHERE
    id = $1
`

// ReadSurvey by id
func (q *Queries) ReadSurvey(ctx context.Context, id uuid.UUID) (SurveysEav, error) {
	row := q.db.QueryRow(ctx, readSurvey, id)
	var i SurveysEav
	err := row.Scan(
		&i.ID,
		&i.DisplayNm,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
