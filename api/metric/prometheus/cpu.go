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
	userTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "me_user_time_us",
		Help: "Mongo Exporter User time in microseconds in the server",
	})

	systemTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "me_system_time_us",
		Help: "Mongo Exporter System time in microseconds in the server",
	})

	totalTime = promauto.NewCounter(prometheus.CounterOpts{
		Name: "me_total_time_us",
		Help: "Mongo Exporter Total time in microseconds in the server",
	})
)

func RecordCpu(ms *service.MetricService) {
	go func() {
		previousTotalTime := 0.0
		previousUserTime := 0.0
		previousSystemTime := 0.0
		for {
			cpu, err := ms.FindCpu(context.Background())
			if err != nil {
				log.Printf("Error retrieving CPU metrics: %v", err)
			} else {
				currentUserTime := float64(cpu.UserTime.Value)
				deltaUserTime := currentUserTime - previousUserTime
				userTime.Add(deltaUserTime)
				previousUserTime = currentUserTime

				currentSystemTime := float64(cpu.SystemTime.Value)
				deltaSystemTime := currentSystemTime - previousSystemTime
				systemTime.Add(deltaSystemTime)
				previousSystemTime = currentSystemTime

				currentTotalTime := float64(cpu.TotalTime.Value)
				deltaTotalTime := currentTotalTime - previousTotalTime
				totalTime.Add(deltaTotalTime)
				previousTotalTime = currentTotalTime
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
