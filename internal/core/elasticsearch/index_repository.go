package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"go-elasticsearch-v1/internal/core"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type IndexRepository struct {
	requestBuilder *elasticsearch.Client
}

func NewIndexRepository(requestBuilder *elasticsearch.Client) *IndexRepository {
	return &IndexRepository{requestBuilder: requestBuilder}
}

func (r *IndexRepository) CreateIndex(ctx context.Context, indexElastic core.Index) (*core.IndexCreateResponse, error) {
	res, err := r.requestBuilder.Indices.Create(
		indexElastic.ID,
		r.requestBuilder.Indices.Create.WithBody(strings.NewReader(fmt.Sprintf("%v", indexElastic.Body))),
	)

	if err != nil {
		return nil, err
	}

	indexJSON := IndexCreateResponseJSON{}
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&indexJSON)

	return r.parseResponse(&indexJSON), nil
}

func (r *IndexRepository) parseResponse(indexJSON *IndexCreateResponseJSON) *core.IndexCreateResponse {
	return &core.IndexCreateResponse{}
}

type IndexCreateResponseJSON struct {
	Acknowledged       bool   `json:"acknowledged"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
	Index              string `json:"index"`
}
