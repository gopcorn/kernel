package sdk

// Request is one call to one operation of one feature.
type Request struct {
	Feature string
	Op      string
	Key     string
	Body    []byte // idempotency key; empty when the caller sent none.
	Caller  Caller
}

// Response is what an operation sends back.
type Response struct {
	Status int    // HTTP style number: 200, 400 etc.
	Body   []byte // the raw bytes that the kernel passes along without validation
}
