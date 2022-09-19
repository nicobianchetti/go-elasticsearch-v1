package create_index

type IndexCreateResult struct {
	Acknowledged       bool
	ShardsAcknowledged bool
	Index              string
}
