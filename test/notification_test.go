package notification_test

import (
	"testing"

	"github.com/jimbertools/beeep/pkg/beeep/notification"
)

func TestNotification(t *testing.T) {
	notification.Notify("Title", "Message", "../test/testdata/information.png")
}
