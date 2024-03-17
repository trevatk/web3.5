package domain

import (
	"context"
	"time"
)

// ValueEnvelope
type ValueEnvelope struct {
	AssessmentID string    `json:"assessment_id"`
	AttributeID  string    `json:"assessment_attribute_id"`
	ValueID      string    `json:"value_id"`
	Value        any       `json:"value"`
	CreatedAt    time.Time `json:"created_at"`
}

// MessageBroker client interface
type MessageBroker interface {
	// Publish message
	Publish(ctx context.Context, topic string, payload []byte) error
	// Close client connection
	Close() error
}
