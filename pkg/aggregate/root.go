package aggregate

type DomainEvent interface {
	GetTopic() string
}

type Root struct {
	events []DomainEvent
}

func (r *Root) AddEvent(event DomainEvent) {
	r.events = append(r.events, event)
}
func (r *Root) PullEvent() []DomainEvent {
	evnts := r.events
	r.events = []DomainEvent{}
	return evnts
}
