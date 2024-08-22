//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !illumos && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!illumos,!js

package notification

import "errors"

var (
	ErrUnsupported = errors.New("unsupported platform")
)

type Notification struct {
	title string
	message string
	iconPath string
}

func NewNotification(title, message, iconPath string) (*Notification, error) {
	return &Notification{title: title, message: message, iconPath: iconPath}, nil
}

func (notification *Notification) Show() error {
	return ErrUnsupported
}

func Notify(Title, Message, iconPath string) error {
	noti, err := NewNotification(Title, Message, iconPath)
	if err != nil {
		return err
	}

	return noti.Show()
}
