package docket

import (
	"eventsourcing/application/fine"
	"eventsourcing/application/user"
)

type DisciplineDocket struct {
	DocketId string
	Name     string
	Users    []user.User
	Fines    []fine.Fine
}
