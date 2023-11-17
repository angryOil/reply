package domain

import (
	"errors"
	"reply/internal/domain/vo"
	"time"
)

var _ Reply = (*reply)(nil)

type Reply interface {
	ValidCreate() error
	ValidUpdate() error

	update(content string) Reply
	ToInfo() vo.Info
}

type reply struct {
	id            int
	cafeId        int
	boardId       int
	writer        int
	content       string
	createdAt     time.Time
	lastUpdatedAt time.Time
}

const (
	InvalidId      = "invalid reply id"
	InvalidWriter  = "invalid writer id"
	InvalidBoard   = "invalid board id"
	InvalidContent = "invalid content"
)

func (r *reply) ValidCreate() error {
	if r.writer < 1 {
		return errors.New(InvalidWriter)
	}
	if r.boardId < 1 {
		return errors.New(InvalidBoard)
	}
	if r.content == "" {
		return errors.New(InvalidContent)
	}
	return nil
}

func (r *reply) ValidUpdate() error {
	if r.id < 1 {
		return errors.New(InvalidId)
	}
	if r.content == "" {
		return errors.New(InvalidContent)
	}
	return nil
}

func (r *reply) ToInfo() vo.Info {
	return vo.Info{
		Id:            r.id,
		CafeId:        r.cafeId,
		BoardId:       r.boardId,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     convertTimeToString(r.createdAt),
		LastUpdatedAt: convertTimeToString(r.lastUpdatedAt),
	}
}

func (r *reply) update(content string) Reply {
	r.content = content
	r.lastUpdatedAt = time.Now()
	return r
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}
