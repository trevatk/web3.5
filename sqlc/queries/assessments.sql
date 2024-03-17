
-- name: InsertAssessmentValue :one
-- InsertAssessmentValue create new assessment value
INSERT INTO assessment_values (
    patient_id,
    assessment_id,
    assessment_attribute_id,
    input
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: ReadAssessmentAttribute :one
-- ReadAssessmentAttribute by id
SELECT
    *
FROM
    assessment_attributes
WHERE
    id = $1
    AND assessment_id = $2;