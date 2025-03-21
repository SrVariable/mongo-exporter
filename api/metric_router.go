package api

import (
	"github.com/SrVariable/mongo-exporter/internal/metric"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func addMetricRoutes(client *mongo.Client, rg *gin.RouterGroup) {
	m := rg.Group("/metrics")

	repo := metric.NewDatabaseRepository(client)
	service := metric.NewMetricService(repo)

	m.GET("/", func(cg *gin.Context) {
		metric.GetMetricsHandler(service)(cg)
	})

	m.GET("/:name", func(cg *gin.Context) {
		metric.GetMetricByNameHandler(service)(cg)
	})
}
