//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package notification

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jimbertools/toast/pkg/toast"
)

type Notification struct {
	toast toast.Toast
}

func (notification *Notification) Show() error {
	return notification.toast.Show()
}

func NewNotification(title, message, iconPath string) (*Notification, error) {
	appId := fmt.Sprintf("com.%s.app", normalizedTitle(title))

	toastManager, err := toast.NewToastManager(appId, title, iconPath)
	if err != nil {
		return nil, err
	}

	toast := toastManager.NewSimpleToast(title, message)
	return &Notification{toast: *toast}, nil
}

func normalizedTitle(title string) string {
	title = strings.ToLower(title)

	re := regexp.MustCompile(`\W`)
	title = re.ReplaceAllString(title, " ")

	re = regexp.MustCompile(`\s+`)
	title = re.ReplaceAllString(title, "_")

	title = strings.Trim(title, "_")

	return title
}

func Notify(Title, Message, iconPath string) error {
	noti, err := NewNotification(Title, Message, iconPath)
	if err != nil {
		return err
	}

	return noti.Show()
}
