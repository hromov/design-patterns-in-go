package main

import (
	"github.com/hromov/design-patterns-in-go/behavioral/observer/publisher"
	"github.com/hromov/design-patterns-in-go/behavioral/observer/subscriber"
)

func main() {
	pub := publisher.New()

	sub1 := subscriber.New("User #1")
	sub2 := subscriber.New("User #2")
	sub3 := subscriber.New("User #3")

	pub.Subscribe(sub1)
	pub.Subscribe(sub2)
	pub.Subscribe(sub3)

	pub.Broadcast("Hello")

	pub.Unsubscribe(sub1)
	pub.Unsubscribe(sub2)
	pub.Unsubscribe(sub3)

	pub.Broadcast("Bye")
}
