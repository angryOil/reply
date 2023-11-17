package model

import (
	"github.com/uptrace/bun"
	"reply/internal/domain"
	"time"
)

type Reply struct {
	bun.BaseModel `bun:"table:board,alias:b"`

	Id            int       `bun:"id,pk,autoincrement"`
	CafeId        int       `bun:"cafe_id,notnull"`
	BoardId       int       `bun:"board_id,notnull"`
	Writer        int       `bun:"writer,notnull"`
	Content       string    `bun:"content,notnull"`
	CreatedAt     time.Time `bun:"created_at"`
	LastUpdatedAt time.Time `bun:"last_updated_at"`
}

func (r Reply) ToDomain() domain.Reply {
	return domain.NewBuilder().
		Id(r.Id).
		CafeId(r.CafeId).
		BoardId(r.BoardId).
		Writer(r.Writer).
		Content(r.Content).
		CreatedAt(r.CreatedAt).
		LastUpdatedAt(r.LastUpdatedAt).
		Build()
}

func ToDomainList(replyList []Reply) []domain.Reply {
	result := make([]domain.Reply, len(replyList))
	for i, r := range replyList {
		result[i] = r.ToDomain()
	}
	return result
}
