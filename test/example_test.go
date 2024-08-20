package beeep

import "github.com/jimbertools/beeep"

func ExampleBeep() {
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
}

func ExampleNotify() {
	beeep.Notify("Title", "MessageBody", "../assets/information.png")
}

func ExampleAlert() {
	beeep.Alert("Title", "MessageBody", "../assets/warning.png")
}
