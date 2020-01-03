package fanout

type staticFanout struct {
	destinations []string
}

func (s *staticFanout) Receivers() ([]string, error) {
	return s.destinations, nil
}

func NewStaticFanout(destinations []string) (Fanout, error) {
	return &staticFanout{
		destinations: destinations,
	}, nil
}
