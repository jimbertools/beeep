//go:build js
// +build js

package notification

import (
	"syscall/js"
	"path/filepath"
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
	defer func() {
		e := recover()

		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
		} else {
			panic(e)
		}
	}()

	n := js.Global().Get("Notification")

	opts := js.Global().Get("Object").Invoke()
	opts.Set("body", notification.message)
	appIcon, err := filepath.Abs(notification.iconPath)
	if err != nil {
		return err
	}
	opts.Set("icon", appIcon)

	if n.Get("permission").String() == "granted" {
		n.New(js.ValueOf(notification.title), opts)
	} else {
		var f js.Func
		f = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if args[0].String() == "granted" {
				n.New(js.ValueOf(notification.title), opts)
			}
			f.Release()
			return nil
		})

		n.Call("requestPermission", f)
	}

	return err
}
