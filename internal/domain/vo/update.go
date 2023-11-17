package vo

import "time"

type Update struct {
	Id            int
	CafeId        int
	BoardId       int
	Writer        int
	Content       string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
