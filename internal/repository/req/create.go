package req

import "time"

type Create struct {
	Writer    int
	CafeId    int
	BoardId   int
	Content   string
	CreatedAt time.Time
}
