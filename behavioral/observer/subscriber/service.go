package subscriber

import "log"

type subscriber struct {
	name string
}

func New(name string) *subscriber {
	return &subscriber{name: name}
}

func (s *subscriber) Notify(msg string) {
	log.Printf("%s: %s", s.name, msg)
}
