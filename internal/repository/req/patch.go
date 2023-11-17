package req

import "time"

type Patch struct {
	Id            int
	CafeId        int
	BoardId       int
	Writer        int
	Content       string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
