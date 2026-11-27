package sdk

import (
	"context"
	"errors"
)

// ErrNotFound is returned by Get when the key does not exist.
// Compare with errors.Is
var ErrNotFound = errors.New("sdk: not found")

// Tx reads and writes values by key.
// Inside Store.InTx, all writes succeed or none do.
type Tx interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
}

// Store is what a feature uses to keep its data.
// It has the same three methods as Tx, plus InTx for grouping several writes.
// example: if the second Put fails, Alice's balance is also not affected.
//
//	 err := store.InTx(ctx, func(tx sdk.Tx) error {
//			if err := tx.Put(ctx, "balance:alice", []byte("400")); err != nil {
//				return err
//			}
//			return tx.Put(ctx, "balance:bob", []byte("600"))
//	 })
type Store interface {
	Tx
	InTx(ctx context.Context, fn func(Tx) error) error
}
