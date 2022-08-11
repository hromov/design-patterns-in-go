package notifier

import "log"

type Notifier interface {
	Send(string)
}

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) Send(msg string) {
	log.Printf("Original message: %s\nWas sent!", msg)
}
