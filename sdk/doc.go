// Package sdk is the contract between the kernel and the features that run on it.

// Two rules to keep the system honest:
//
// 1. The package imports nothing from this repository.
// 2. A feature imports only this package (sdk) and other features' api modules.
//
// A feature implements Feature, registers itself using Register from an init function,
// and declares its operation in Init with Host.Handle
package sdk
