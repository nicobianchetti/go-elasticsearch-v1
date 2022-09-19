package main

import (
	"go-elasticsearch-v1/cmd/api/routes"
	"go-elasticsearch-v1/cmd/api/routes/handlers"
	repositoryES "go-elasticsearch-v1/internal/core/elasticsearch"
	platformES "go-elasticsearch-v1/internal/platform/elasticsearch"
	"go-elasticsearch-v1/internal/platform/logs"
	"go-elasticsearch-v1/internal/usecase/create_index"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	EnvPort = "port"
	Port    = "8080"

	msgElasticsearchFail = "[event: elasticsearch_init] Elasticsearch init failed"
	msgAppRunFail        = "[event: fail_api_run] App run error"
)

func main() {
	router := injectDependencies()
	// Start serving the application
	port := os.Getenv(EnvPort)
	if port == "" {
		port = Port
	}

	if err := router.Run(":" + port); err != nil {
		logs.Panic(msgAppRunFail, err)
	}
}

func injectDependencies() *gin.Engine {
	//Client ES
	elasticSearchClient, err := platformES.NewClient()
	if err != nil {
		logs.Panic(msgElasticsearchFail, err)
	}

	//Repositories
	indexElasticSearchRepository := repositoryES.NewIndexRepository(elasticSearchClient)

	//Service
	indexService := create_index.NewService(indexElasticSearchRepository)

	//Handler
	indexHandler := handlers.NewPutIndexHandler(indexService)

	//Routes
	router := gin.Default()
	// router.GET(routes.GetDocuments, )
	// router.POST(routes.PostDocuments, )
	router.PUT(routes.PutIndex, indexHandler.Handle)

	return router
}
