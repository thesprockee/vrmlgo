package vrmlgo

// EventHandler is an interface for Discord events.
type EventHandler interface {
	// Type returns the type of event this handler belongs to.
	Type() string

	// Handle is called whenever an event of Type() happens.
	// It is the receivers responsibility to type assert that the interface
	// is the expected struct.
	Handle(*Session, interface{})
}

// EventInterfaceProvider is an interface for providing empty interfaces for
// Discord events.
type EventInterfaceProvider interface {
	// Type is the type of event this handler belongs to.
	Type() string

	// New returns a new instance of the struct this event handler handles.
	// This is called once per event.
	// The struct is provided to all handlers of the same Type().
	New() interface{}
}

// interfaceEventType is the event handler type for interface{} events.
const interfaceEventType = "__INTERFACE__"

// interfaceEventHandler is an event handler for interface{} events.
type interfaceEventHandler func(*Session, interface{})

// Type returns the event type for interface{} events.
func (eh interfaceEventHandler) Type() string {
	return interfaceEventType
}

// Handle is the handler for an interface{} event.
func (eh interfaceEventHandler) Handle(s *Session, i interface{}) {
	eh(s, i)
}

var registeredInterfaceProviders = map[string]EventInterfaceProvider{}

// registerInterfaceProvider registers a provider so that DiscordGo can
// access it's New() method.
func registerInterfaceProvider(eh EventInterfaceProvider) {
	if _, ok := registeredInterfaceProviders[eh.Type()]; ok {
		return
		// XXX:
		// if we should error here, we need to do something with it.
		// fmt.Errorf("event %s already registered", eh.Type())
	}
	registeredInterfaceProviders[eh.Type()] = eh
	return
}

// eventHandlerInstance is a wrapper around an event handler, as functions
// cannot be compared directly.
type eventHandlerInstance struct {
	eventHandler EventHandler
}
