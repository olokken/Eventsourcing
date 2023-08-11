package application

type Aggregate struct {
	Events []IEvent
}

func (a *Aggregate) AddEvent(e IEvent) {
	a.Events = append(a.Events, e)
}
