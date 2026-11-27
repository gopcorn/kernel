// sdk/bus.go: telling other features that something has happened.
package sdk

import "context"

// Event is a message published to a topic.
type Event struct {
	Topic   string
	Payload []byte
}

// EventHandler reacts to an event.
// Return an error and the event may be retried.
// Write handlers so that running one twice does no harm.
type EventHandler func(ctx context.Context, e Event) error

// Bus publishes events and lets a feature subscribe to a topic.
// a feature calls Publish to announce something, others Subscribe when initializing to listen to
// such events.
type Bus interface {
	Publish(ctx context.Context, e Event) error
	Subscribe(topic string, h EventHandler)
}
