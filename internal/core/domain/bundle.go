package domain

import "context"

// Bundler
type Bundler interface {
	Ping(ctx context.Context) error
}
