package prometheus

import (
	"context"
	"log"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	current = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_current_connections",
		Help: "Mongo Exporter Current connections in the server",
	})
	available = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_available_connections",
		Help: "Mongo Exporter Available connections in the server",
	})
	totalCreated = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_total_created_connections",
		Help: "Mongo Exporter Total connections created in the server",
	})
	active = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_active_connections",
		Help: "Mongo Exporter Active connections in the server",
	})
)

func RecordConnections(ms *service.MetricService) {
	go func() {
		for {
			connections, err := ms.FindConnections(context.Background())
			if err != nil {
				log.Printf("Error retrieving RAM metrics: %v", err)
			} else {
				current.Set(float64(connections.Current.Value))
				available.Set(float64(connections.Available.Value))
				totalCreated.Set(float64(connections.TotalCreated.Value))
				active.Set(float64(connections.Active.Value))
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
