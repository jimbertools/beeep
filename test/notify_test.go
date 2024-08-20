package beeep

import (
	"testing"

	"github.com/jimbertools/beeep"
)

func TestNotify(t *testing.T) {
	err :=	beeep.Notify("Notify title", "Message body", "../assets/information.png")
	if err != nil {
		t.Error(err)
	}
}
