package application

import "time"

type IEvent interface {
	IsEvent()
}

type Event struct {
	_id           string
	AggregateType AggregateType
	Version       int
	Timestamp     time.Time
}

type AggregateType int

const (
	User AggregateType = iota
	DisciplineDocket
	Allocation
	Fine
)

func CreateEvent(aggregateType AggregateType, version int) Event {
	return Event{
		AggregateType: aggregateType,
		Timestamp:     time.Now(),
		Version:       version,
	}
}
