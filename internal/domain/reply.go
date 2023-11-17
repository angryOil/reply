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

	Update(content string) (Reply, error)
	ToDetail() vo.Detail
	ToInfo() vo.Info
	ToUpdate() vo.Update
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

func (r *reply) ToUpdate() vo.Update {
	return vo.Update{
		Id:            r.id,
		CafeId:        r.cafeId,
		BoardId:       r.boardId,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     r.createdAt,
		LastUpdatedAt: r.lastUpdatedAt,
	}
}

func (r *reply) ToDetail() vo.Detail {
	return vo.Detail{
		Id:            r.id,
		CafeId:        r.cafeId,
		BoardId:       r.boardId,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     convertTimeToString(r.createdAt),
		LastUpdatedAt: convertTimeToString(r.lastUpdatedAt),
	}
}

func (r *reply) ToInfo() vo.Info {
	return vo.Info{
		Id:            r.id,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     convertTimeToString(r.createdAt),
		LastUpdatedAt: convertTimeToString(r.lastUpdatedAt),
	}
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

func (r *reply) Update(content string) (Reply, error) {
	r.content = content
	r.lastUpdatedAt = time.Now()
	err := r.ValidUpdate()
	if err != nil {
		return r, err
	}
	return r, nil
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}
