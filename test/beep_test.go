package beeep

import (
	"testing"

	"github.com/jimbertools/beeep"
)

func TestBeep(t *testing.T) {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		t.Error(err)
	}
}
