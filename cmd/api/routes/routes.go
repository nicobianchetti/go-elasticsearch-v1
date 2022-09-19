package routes

const (
	basePath = "go-elasticsearch-v1"

	PostDocuments = basePath + "/doc/:index_id"

	GetDocuments = basePath + "/search/:index_id"

	PutIndex = basePath + "/index"
)
