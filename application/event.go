package application

import (
	"time"

	"github.com/google/uuid"
)

type IEvent interface {
	IsEvent()
}

type Event struct {
	EventId       string
	AggregateType AggregateType
	EventType     EventType
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

type EventType int

const (
	UserCreated EventType = iota
	NameChanged
)

func CreateEvent(aggregateType AggregateType, eventType EventType, version int) Event {
	return Event{
		EventId:       uuid.New().String(),
		AggregateType: aggregateType,
		Timestamp:     time.Now(),
		Version:       version,
		EventType:     eventType,
	}
}
