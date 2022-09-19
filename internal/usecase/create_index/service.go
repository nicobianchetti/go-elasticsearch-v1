package create_index

import (
	"context"
	"go-elasticsearch-v1/internal/core"
)

type IndexWritter interface {
	CreateIndex(ctx context.Context, indexElastic core.Index) (*core.IndexCreateResponse, error)
}

type defaultService struct {
	indexRepository IndexWritter
}

type Params struct {
	IndexID string
	Body    interface{}
}

func NewService(indexRepository IndexWritter) *defaultService {
	return &defaultService{indexRepository: indexRepository}
}

func (s *defaultService) CreateIndex(ctx context.Context, params Params) (*IndexCreateResult, error) {
	res, err := s.indexRepository.CreateIndex(ctx, core.Index{ID: params.IndexID, Body: params.Body})
	if err != nil {
		return nil, err
	}

	return &IndexCreateResult{
		Acknowledged:       res.Acknowledged,
		ShardsAcknowledged: res.ShardsAcknowledged,
		Index:              res.Index,
	}, nil
}
