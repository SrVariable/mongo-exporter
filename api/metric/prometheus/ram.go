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
	resident = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_resident_mib",
		Help: "Mongo Exporter Resident memory (MiB) used in the server",
	})

	virtual = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "me_virtual_mib",
		Help: "Mongo Exporter Virtual memory (MiB) used in the server",
	})
)

// TODO: Calculate the usage of RAM instead of showing resident/virtual RAM
func RecordRam(ms *service.MetricService) {
	go func() {
		for {
			ram, err := ms.FindRam(context.Background())
			if err != nil {
				log.Printf("Error retrieving RAM metrics: %v", err)
			} else {
				resident.Set(float64(ram.Resident.Value))
				virtual.Set(float64(ram.Virtual.Value))
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
