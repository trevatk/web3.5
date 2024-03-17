package controller

var (
	assessmentsPath   = "/assessments"
	assessmentsIDPath = "/assessments/:assessment_id"

	assessmentsAttributePath   = assessmentsIDPath + "/attributes"
	assessmentsAttributeIDPath = assessmentsAttributePath + "/:attribute_id"

	assessmentAttributeValuePath = assessmentsAttributeIDPath + "/value"
)
