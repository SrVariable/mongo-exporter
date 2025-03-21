package metric

import (
	"context"
)

type MetricService interface {
	FindMetrics(c context.Context) ([]Metric, error)
	FindMetricByName(c context.Context, name string) (*Metric, error)
}

type metricService struct {
	repo MetricRepository
}

func NewMetricService(repo MetricRepository) *metricService {
	return &metricService{repo}
}

func (ms *metricService) FindMetrics(c context.Context) ([]Metric, error){
	return ms.repo.GetMetrics(c)
}

func (ms *metricService) FindMetricByName(c context.Context, name string) (*Metric, error) {
	return ms.repo.GetMetricByName(c, name)
}
