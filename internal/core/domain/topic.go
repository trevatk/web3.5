package domain

// Topic
type Topic string

const (
	// CreateAssessmentAttributeValue
	CreateAssessmentAttributeValue Topic = "assessment_attribute_value.create"
)

// String
func (t Topic) String() string {
	switch t {
	case CreateAssessmentAttributeValue:
		return string(CreateAssessmentAttributeValue)
	default:
		return ""
	}
}
