//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package notification

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"syscall"

	"github.com/go-toast/toast"
)

type Windows10Notification struct {
	notification toast.Notification
}

func (w *Windows10Notification) Show() error {
	return w.notification.Push()
}

func NewWindows10Toast(title, message, iconPath string) Notification {
	windows10ToastAppID := windows10NotificationAppID()
	return &Windows10Notification{
		notification: toast.Notification{
			AppID:   windows10ToastAppID,
			Title:   title,
			Message: message,
			Icon:    iconPath,
		},
	}
}

func windows10NotificationAppID() string {
	defID := "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe"
	cmd := exec.Command("powershell", "-NoProfile", "Get-StartApps")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return defID
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "powershell.exe") {
			sp := strings.Split(line, " ")
			if len(sp) > 0 {
				return sp[len(sp)-1]
			}
		}
	}

	return defID
}
