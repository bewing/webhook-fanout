package fanout

import (
	"math/rand"
	"time"
)

// randomFanout is an implemetation of Fanout that generates random strings for fanouts
// for testing purposes
type randomFanout struct {
	sources []string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (f *randomFanout) Receivers() ([]string, error) {
	return f.sources, nil
}

var charset = []rune("abcdefghijklmnopqrstuvwxyz")

func (f *randomFanout) generateFanout() (string, error) {
	r := make([]rune, 8)
	for i := range r {
		r[i] = charset[rand.Intn(len(charset))]
	}
	return string(r), nil
}

// NewRandomFanout creates a new randomFanout with random targets
func NewRandomFanout() (Fanout, error) {
	f := &randomFanout{}
	for i := 0; i < 8; i++ {
		newFanout, _ := f.generateFanout()
		f.sources = append(f.sources, newFanout)
	}
	return f, nil
}
