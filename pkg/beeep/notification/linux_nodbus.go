//go:build (linux && nodbus) || (freebsd && nodbus) || (netbsd && nodbus) || (openbsd && nodbus) || illumos
// +build linux,nodbus freebsd,nodbus netbsd,nodbus openbsd,nodbus illumos

package notification

import (
	"errors"
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
	
	cmd := func() error {
		send, err := exec.LookPath("sw-notify-send")
		if err != nil {
			send, err = exec.LookPath("notify-send")
			if err != nil {
				return err
			}
		}

		c := exec.Command(send, notification.title, notification.message, "-i", notification.iconPath)
		return c.Run()
	}

	knotify := func() error {
		send, err := exec.LookPath("kdialog")
		if err != nil {
			return err
		}
		c := exec.Command(send, "--title", notification.title, "--passivepopup", notification.message, "10", "--icon", notification.iconPath)
		return c.Run()
	}

	err := cmd()
	if err != nil {
		e := knotify()
		if e != nil {
			return errors.New("beeep: " + err.Error() + "; " + e.Error())
		}
	}

	return nil
}