
-- name: InsertSurvey :one
-- InsertSurvey insert new survey record
INSERT INTO surveys_eav (
    display_nm,
    description
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ReadSurvey :one
-- ReadSurvey by id
SELECT  
    *
FROM
    surveys_eav
WHERE
    id = $1;

-- name: ListSurveys :many
-- ListSurveys retrieve partial survey records
SELECT
    id,
    display_nm
FROM
    surveys_eav
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: InsertSurveyAttribute :one
-- InsertSurveyAttribute 
INSERT INTO survey_attributes (
    survey_id,
    display_nm,
    description,
    dtype
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListSurveyAttributes :many
-- ListSurbeyAttributes 
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
ORDER BY order_execution ASC;

-- name: InsertSurveyValue :one
-- InsertSurveyValue insert new survey value record
INSERT INTO survey_values (
    submitted_by,
    survey_id,
    survey_attribute_id,
    input
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;