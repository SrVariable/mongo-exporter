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
	insert = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_insert_total_opcounter",
		Help: "Mongo Exporter Total number of insert operations in the server",
	})
	delete = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_delete_total_opcounter",
		Help: "Mongo Exporter Total number of delete operations in the server",
	})
	query = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_query_total_opcounter",
		Help: "Mongo Exporter Total number of query operations in the server",
	})
	update = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_update_total_opcounter",
		Help: "Mongo Exporter Total number of update operations in the server",
	})
)

func RecordOpCounters(ms *service.MetricService) {
	go func() {
		for {
			opcounters, err := ms.FindOpCounters(context.Background())
			if err != nil {
				log.Printf("Error retrieving OpCounters metrics: %v", err)
				continue
			} else {
				insert.Set(float64(opcounters.Insert.Value))
				delete.Set(float64(opcounters.Delete.Value))
				update.Set(float64(opcounters.Update.Value))
				query.Set(float64(opcounters.Query.Value))
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
