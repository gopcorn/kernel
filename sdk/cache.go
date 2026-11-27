package sdk

import "context"

// Cache keeps short lived values.
// a "space" is a named group of keys with its own expiry time,
// declared by the feature (phase 9 adds the declaring part).
type Cache interface {
	Get(ctx context.Context, space, key string) (value []byte, ok bool, err error)
	Set(ctx context.Context, space, key string, value []byte) error
	Delete(ctx context.Context, space, key string) error
}
