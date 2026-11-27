// sdk/op.go : an operation that a feature can do.
package sdk

import "context"

// Handler runs one operation.
// It receives the request and returns a response.
type Handler func(ctx context.Context, req Request) Response

// Op describes one operation that a feature offers
type Op struct {
	Name     string
	Access   Access
	Request  any
	Response any
	Handle   Handler
}
