//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package notification

type WindowsVersion int

const (
	Windows10 WindowsVersion = iota + 10
	Windows11
)

type Notification interface {
	Show() error
}
