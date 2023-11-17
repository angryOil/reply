package controller

import (
	"context"
	"reply/internal/controller/req"
	"reply/internal/controller/res"
	"reply/internal/page"
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

func (c Controller) GetList(ctx context.Context, boardId int, reqPage page.ReqPage) ([]res.GetList, int, error) {
	list, total, err := c.s.GetListTotal(ctx, boardId, reqPage)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(list))
	for i, l := range list {
		dto[i] = res.GetList{
			Id:            l.Id,
			Writer:        l.Writer,
			Content:       l.Content,
			CreatedAt:     l.CreatedAt,
			LastUpdatedAt: l.LastUpdatedAt,
		}
	}
	return dto, total, nil
}

func (c Controller) GetCountList(ctx context.Context, idArr []int) ([]res.GetCountList, error) {
	list, err := c.s.GetCountList(ctx, idArr)
	if err != nil {
		return []res.GetCountList{}, err
	}
	dto := make([]res.GetCountList, len(list))
	for i, l := range list {
		dto[i] = res.GetCountList{
			BoardId:    l.BoardId,
			ReplyCount: l.Count,
		}
	}
	return dto, nil
}
