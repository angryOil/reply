package service

import (
	"context"
	"errors"
	"reply/internal/domain"
	"reply/internal/page"
	"reply/internal/repository"
	"reply/internal/repository/req"
	req2 "reply/internal/service/req"
	"reply/internal/service/res"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

const (
	NoRows = "no rows"
)

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

func (s Service) Patch(ctx context.Context, p req2.Patch) error {
	id := p.Id
	content := p.Content
	err := domain.NewBuilder().
		Id(id).
		Content(content).
		Build().ValidUpdate()
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, id,
		func(domains []domain.Reply) (domain.Reply, error) {
			if len(domains) != 1 {
				return nil, errors.New(NoRows)
			}
			return domains[0], nil
		},
		func(d domain.Reply) (req.Patch, error) {
			u, err := d.Update(content)
			if err != nil {
				return req.Patch{}, err
			}
			v := u.ToUpdate()
			return req.Patch{
				Id:            v.Id,
				CafeId:        v.CafeId,
				BoardId:       v.BoardId,
				Writer:        v.Writer,
				Content:       v.Content,
				CreatedAt:     v.CreatedAt,
				LastUpdatedAt: v.LastUpdatedAt,
			}, nil
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	return err
}

func (s Service) GetListTotal(ctx context.Context, boardId int, reqPage page.ReqPage) ([]res.GetListTotal, int, error) {
	domains, total, err := s.repo.GetListTotal(ctx, boardId, reqPage)
	if err != nil {
		return []res.GetListTotal{}, 0, err
	}
	result := make([]res.GetListTotal, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		result[i] = res.GetListTotal{
			Id:            v.Id,
			Writer:        v.Writer,
			Content:       v.Content,
			CreatedAt:     v.CreatedAt,
			LastUpdatedAt: v.LastUpdatedAt,
		}
	}
	return result, total, nil
}

func (s Service) GetCountList(ctx context.Context, arr []int) ([]res.GetCountList, error) {
	list, err := s.repo.GetCountList(ctx, arr)
	if err != nil {
		return []res.GetCountList{}, err
	}
	dto := make([]res.GetCountList, len(list))
	for i, l := range list {
		dto[i] = res.GetCountList{
			BoardId: l.BoardId,
			Count:   l.Count,
		}
	}
	return dto, err
}
