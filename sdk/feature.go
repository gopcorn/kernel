// sdk/feature.go: the two interfaces that a feature must meet.
package sdk

import (
	"context"
	"log/slog"
)

// Feature is what you write. The kernel calls Init once at startup.
type Feature interface {
	Name() string
	Init(h Host) error
}

// Host is everything a feature is allowed to use.
// If it is not on this list, that feature cannot touch it.
type Host interface {
	// Handle declares an operation. Call it once per operation, from Init.
	Handle(op Op)
	// Call runs an operation of another feature through the kernel.
	Call(ctx context.Context, feature, op, key string, body []byte) (Response, error)
	Store() Store
	Cache() Cache
	Bus() Bus
	// the kernel will return a slog.Logger that already prints the feature name.
	Log() *slog.Logger
}
