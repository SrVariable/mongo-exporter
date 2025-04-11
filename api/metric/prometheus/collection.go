package prometheus

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SrVariable/mongo-exporter/internal/metric/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// NOTE: To collect metrics from a certain collection, replace dbName and collName accordingly
var (
	dbName     = "test"
	collName   = "users"
	collInsert = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("me_coll_%s_%s_inserts", dbName, collName),
		Help: fmt.Sprintf("Mongo Exporter Number of inserts in %s.%s", dbName, collName),
	})
	collRemove = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("me_coll_%s_%s_removes", dbName, collName),
		Help: fmt.Sprintf("Mongo Exporter Number of removes in %s.%s", dbName, collName),
	})
	collQueries = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("me_coll_%s_%s_queries", dbName, collName),
		Help: fmt.Sprintf("Mongo Exporter Number of queries in %s.%s", dbName, collName),
	})
	collUpdate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("me_coll_%s_%s_updates", dbName, collName),
		Help: fmt.Sprintf("Mongo Exporter Number of updates in %s.%s", dbName, collName),
	})
)

func RecordCollection(ms *service.MetricService) {
	go func() {
		for {
			collection, err := ms.FindCollection(context.Background(), dbName, collName)
			if err != nil {
				log.Printf("Error retrieving Collection metrics: %v", err)
			} else {
				collInsert.Set(float64(collection.Insert.Value))
				collRemove.Set(float64(collection.Remove.Value))
				collUpdate.Set(float64(collection.Update.Value))
				collQueries.Set(float64(collection.Queries.Value))
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
