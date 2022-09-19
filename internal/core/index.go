package core

type Index struct {
	ID   string
	Body interface{}
}

type IndexCreateResponse struct {
	Acknowledged       bool
	ShardsAcknowledged bool
	Index              string
}
