//go:build darwin && !linux && !freebsd && !netbsd && !openbsd && !windows && !js
// +build darwin,!linux,!freebsd,!netbsd,!openbsd,!windows,!js

package notification

import (
	"fmt"
	"os/exec"
)

type Notification struct {
	title    string
	message  string
	iconPath string
}

func NewNotification(title, message, iconPath string) (*Notification, error) {
	return &Notification{title: title, message: message, iconPath: iconPath}, nil
}

func (notification *Notification) Show() error {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return err
	}

	script := fmt.Sprintf("display notification %q with title %q", notification.message, notification.title)
	cmd := exec.Command(osa, "-e", script)
	return cmd.Run()
}

func Notify(Title, Message, iconPath string) error {
	noti, err := NewNotification(Title, Message, iconPath)
	if err != nil {
		return err
	}

	return noti.Show()
}
