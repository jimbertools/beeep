package beeep

import (
	"testing"

	"github.com/jimbertools/beeep"
)

func TestAlert(t *testing.T) {
	err := beeep.Alert("Alert title", "Message body", "../assets/warning.png")
	if err != nil {
		t.Error(err)
	}
}
