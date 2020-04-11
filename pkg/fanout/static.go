package fanout

type staticFanout struct {
	destinations []string
}

func (s *staticFanout) Receivers() ([]string, error) {
	return s.destinations, nil
}

// NewStaticFanout returns a fanout to a static list of strings
func NewStaticFanout(destinations []string) (Fanout, error) {
	return &staticFanout{
		destinations: destinations,
	}, nil
}
