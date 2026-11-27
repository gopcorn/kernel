package sdk

// CallerKind says what kind of caller made the request.
type CallerKind int

const (
	// Anonymous: no sign-in at all.
	Anonymous CallerKind = iota // Anonymous is 0 on purpose.
	// EndUser: a signed-in customer
	EndUser
	// Operator: a signed in staff member
	Operator
	// FromFeature: another feature calling through the host.
	FromFeature
)

// Caller identifies who is making a request.
// the kernel does not prove the identity.
// the authentication function does this and fills in.
type Caller struct {
	Kind    CallerKind
	Subject string // the user/ operator ID. empty when anonymous.
	Feature string // the name of the calling feature. Only set when Kind is FromFeature.
}
