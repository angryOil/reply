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

func (c Controller) Patch(ctx context.Context, id int, p req.Patch) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:      id,
		Content: p.Content,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, id int) error {
	err := c.s.Delete(ctx, id)
	return err
}
