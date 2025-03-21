package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMetricsHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, err := ms.FindMetrics(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}

func GetMetricByNameHandler(ms MetricService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		m, err := ms.FindMetricByName(c.Request.Context(), name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, m)
	}
}
