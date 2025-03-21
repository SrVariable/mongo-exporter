package metric

import (
	"fmt"
	"errors"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MetricRepository interface {
	GetMetrics(c context.Context) ([]Metric, error)
	GetMetricByName(c context.Context, name string) (*Metric, error)
}

type DatabaseRepository struct {
	Client *mongo.Client
}

func NewDatabaseRepository(client *mongo.Client) *DatabaseRepository {
	return &DatabaseRepository{
		Client: client,
	}
}

// It currently shows `opcounters`, but can be modified
// depending on the specific metrics we want to get
func (dr *DatabaseRepository) GetMetrics(c context.Context) ([]Metric, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	if err != nil {
		return nil, err
	}

	var metrics []Metric

	opcounters, ok := result["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}
	for metricName, metricValue := range opcounters {
		now := time.Now()
		metric := Metric{
			Name: metricName,
			Value: fmt.Sprintf("%d", metricValue),
			Timestamp: fmt.Sprintf(
				"%4d-%02d-%02d %02d:%02d:%02d",
				now.Year(), now.Month(), now.Day(),
				now.Hour(), now.Minute(), now.Second(),
			),
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (dr *DatabaseRepository) GetMetricByName(c context.Context, name string) (*Metric, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	if err != nil {
		return nil, err
	}

	opcounters, ok := result["opcounters"].(bson.M)
	if !ok {
		return nil, errors.New("wrong type")
	}

	value, ok := opcounters[name]
	if !ok {
		return nil, errors.New("metric not found")
	}

	now := time.Now()
	metric := Metric{
		Name: name,
		Value: fmt.Sprintf("%d", value),
		Timestamp: fmt.Sprintf(
			"%4d-%02d-%02d %02d:%02d:%02d",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second(),
		),
	}
	return &metric, nil
}
