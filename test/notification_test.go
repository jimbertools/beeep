package notification_test

import (
	"testing"

	"github.com/jimbertools/beeep/pkg/beeep/notification"
)

func TestNotification(t *testing.T) {

	notification, err := notification.NewNotification("Title", "Message", "../test/testdata/information.png")
	if err != nil {
		t.Error(err)
	}

	err = notification.Show()
	if err != nil {
		t.Error(err)
	}
}
