package allocation

import "time"

type Allocation struct {
	AllocationId string
	UserId       string
	FineId       string
	Date         time.Time
}
