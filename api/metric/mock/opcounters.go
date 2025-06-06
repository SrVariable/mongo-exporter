package service

import (
	"net/http"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/gin-gonic/gin"
)

func GetOpCountersHandlerMock(ms *service.MetricService, c *gin.Context) {
	opcounters, err := ms.FindOpCounters(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, opcounters)
}
