package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics
var (
	RequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
		Help: "The total number of requests",
	})
)
