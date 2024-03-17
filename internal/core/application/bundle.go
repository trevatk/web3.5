package application

import "context"

// Bundle
type Bundle struct{}

// Ping
func (bu *Bundle) Ping(ctx context.Context) error {
	return nil
}
