//go:build (linux && !nodbus) || (freebsd && !nodbus) || (netbsd && !nodbus) || (openbsd && !nodbus)
// +build linux,!nodbus freebsd,!nodbus netbsd,!nodbus openbsd,!nodbus

package notification

import (
	"errors"
	"os/exec"
	"path/filepath"

	"github.com/godbus/dbus/v5"
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
	appIcon, err := filepath.Abs(notification.iconPath)
	if err != nil {
		return err
	}

	cmd := func() error {
		send, err := exec.LookPath("sw-notify-send")
		if err != nil {
			send, err = exec.LookPath("notify-send")
			if err != nil {
				return err
			}
		}

		c := exec.Command(send, notification.title,	notification.message, "-i", appIcon)
		return c.Run()
	}

	knotify := func() error {
		send, err := exec.LookPath("kdialog")
		if err != nil {
			return err
		}
		c := exec.Command(send, "--title", notification.title, "--passivepopup", notification.message, "10", "--icon", appIcon)
		return c.Run()
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		return cmd()
	}
	obj := conn.Object("org.freedesktop.Notifications", dbus.ObjectPath("/org/freedesktop/Notifications"))

	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0), appIcon, notification.title, notification.message, []string{}, map[string]dbus.Variant{}, int32(-1))
	if call.Err != nil {
		e := cmd()
		if e != nil {
			e := knotify()
			if e != nil {
				return errors.New("beeep: " + call.Err.Error() + "; " + e.Error())
			}
		}
	}

	return nil
}