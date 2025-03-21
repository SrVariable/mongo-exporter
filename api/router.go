package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var router = gin.Default()

func Run(client *mongo.Client, uri string) (err error) {
	router := setupRouter(client)
	return router.Run(uri)
}

func setupRouter(client *mongo.Client) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/v1")

	addHealthcheckRoutes(v1)
	addHelloRoutes(v1)
	addMetricRoutes(client, v1)
	addStudentRoutes(v1)

	return router
}
