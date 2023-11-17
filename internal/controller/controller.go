package controller

import (
	"context"
	"reply/internal/controller/req"
	"reply/internal/service"
	req2 "reply/internal/service/req"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, cafeId int, boardId int, create req.Create) error {
	err := c.s.Create(ctx, req2.Create{
		Writer:  create.Writer,
		CafeId:  cafeId,
		BoardId: boardId,
		Content: create.Content,
	})
	return err
}
