//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !illumos && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!illumos,!js

package notification

import "errors"

var (
	ErrUnsupported = errors.New("unsupported platform")
)

// Notify sends desktop notification.
func Notify(title, message, appIcon string) error {
	return ErrUnsupported
}
