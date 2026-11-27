// sdk/contract_test.go

package sdk_test // different package intentionally

import (
	"context"
	"testing"

	"github.com/gopcorn/kernel/sdk"
)

// noop is a throwaway feature.
// It lives in package sdk_test, an *outside*
// package, so it can use only what the sdk exports.
// If this file compiles, a
// real feature can do everything it needs with the contract.
type noop struct{}

func (noop) Name() string { return "noop" }

func (noop) Init(h sdk.Host) error {
	h.Handle(sdk.Op{
		Name:   "ping",
		Access: sdk.Public,
		Handle: func(ctx context.Context, r sdk.Request) sdk.Response {
			h.Log().Info("ping", "caller", r.Caller.Kind)
			return sdk.Response{Status: 200, Body: []byte("pong")}
		},
	})
	return nil
}

func TestNoopIsAFeature(t *testing.T) {
	var f sdk.Feature = noop{}
	if f.Name() != "noop" {
		t.Fatalf("name = %q", f.Name())
	}
}
