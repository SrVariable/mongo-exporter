package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SrVariable/mongo-exporter/internal/metric/service"
)

func GetOpCountersHandler(ms service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindOpCounters(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}

func GetOpCounterByNameHandler(ms service.MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		m, err := ms.FindOpCounterByName(c.Request.Context(), name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
