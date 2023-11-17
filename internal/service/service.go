package service

import (
	"context"
	"reply/internal/domain"
	"reply/internal/repository"
	"reply/internal/repository/req"
	req2 "reply/internal/service/req"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, c req2.Create) error {
	writer, cafeId, boardId := c.Writer, c.BoardId, c.CafeId
	content := c.Content
	createdAt := time.Now()
	err := domain.NewBuilder().
		Writer(writer).
		CafeId(cafeId).
		BoardId(boardId).
		Content(content).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req.Create{
		Writer:    writer,
		CafeId:    cafeId,
		BoardId:   boardId,
		Content:   content,
		CreatedAt: createdAt,
	})
	return err
}
