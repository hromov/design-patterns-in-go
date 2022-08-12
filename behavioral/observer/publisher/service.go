package publisher

type Subscriber interface {
	Notify(string)
}

type Publisher interface {
	Subscribe(Subscriber)
	Unsubscribe(Subscriber)
	Broadcast(string)
}

type publisher struct {
	subscribers map[Subscriber]struct{}
}

func New() *publisher {
	return &publisher{subscribers: map[Subscriber]struct{}{}}
}

func (p *publisher) Subscribe(s Subscriber) {
	p.subscribers[s] = struct{}{}
}

func (p *publisher) Unsubscribe(s Subscriber) {
	if _, ok := p.subscribers[s]; ok {
		delete(p.subscribers, s)
	}
}

func (p *publisher) Broadcast(msg string) {
	for s := range p.subscribers {
		s.Notify(msg)
	}
}
