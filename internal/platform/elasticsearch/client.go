package elasticsearch

import (
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	uri = "http://localhost:9200"
)

func NewClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{uri},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		cause := fmt.Sprintf("Failed to connect to Elasticsearch: %s", err.Error())
		return nil, errors.New(cause)
	}

	return es, nil
}
