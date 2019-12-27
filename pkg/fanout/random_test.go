package fanout

import (
	"testing"
)

func TestRandomFanout(t *testing.T) {
	f, _ := NewRandomFanout()
	receivers, _ := f.Receivers()
	count := len(receivers)
	if count != 8 {
		t.Error(count)
	}
}
