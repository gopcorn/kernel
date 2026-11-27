package sdk

import "fmt"

// Access says who may call an operation.
// enums below define levels of access.
type Access int

const (
	// Public: anyone, signed in or not., 1=Public, 2=User
	Public Access = iota + 1
	// User: a signed-in end user, or an operator.
	User
	// Admin: an operator only
	Admin
	// Internal: other features only. Outside callers never see it.
	Internal
)

// String makes it easy to identify the access level e.g. in logs.
func (a Access) String() string {
	switch a {
	case Public:
		return "public"
	case User:
		return "user"
	case Admin:
		return "admin"
	case Internal:
		return "internal"
	}
	return fmt.Sprintf("access(%d)", int(a))
}
