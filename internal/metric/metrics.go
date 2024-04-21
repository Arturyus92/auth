package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "my_space"
	appName   = "auth_service"
)

// Metrics...
type Metrics struct {
	requestCounter prometheus.Counter
	//responseCounter       *prometheus.CounterVec
	//histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

// Init...
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_requests_total",
				Help:      "Количество запросов к серверу",
			},
		),
	}
	return nil
}

// IncRequestCounter...
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}
