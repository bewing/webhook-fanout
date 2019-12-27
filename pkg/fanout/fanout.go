package fanout

// Fanout defines the interface for returning target receivers.
type Fanout interface {
	Receivers() ([]string, error)
}
