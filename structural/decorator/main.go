package main

import (
	fbnotifier "github.com/hromov/design-patterns-in-go/structural/decorator/fbNotifier"
	"github.com/hromov/design-patterns-in-go/structural/decorator/notifier"
)

func main() {
	original := notifier.New()
	fbNotifier := fbnotifier.New(original)

	fbNotifier.Send("this is the message")
}
