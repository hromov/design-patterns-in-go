package fbnotifier

import (
	"log"

	"github.com/hromov/design-patterns-in-go/structural/decorator/notifier"
)

type service struct {
	n notifier.Notifier
}

func New(n notifier.Notifier) *service {
	return &service{n}
}

func (s *service) Send(msg string) {
	s.n.Send(msg)
	log.Printf("Msg: %s\nWas sent to FB", msg)
}
