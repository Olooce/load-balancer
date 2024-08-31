package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Shared metric for request counting across servers
var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "load_balancer_requests_total",
			Help: "Total number of requests received by the load balancer",
		},
		[]string{"server"}, // Label by server
	)
)

// Initialize and register Prometheus metrics
func init() {
	prometheus.MustRegister(RequestCount)
}

// Increment the request counter for a specific server
func RecordRequest(server string) {
	RequestCount.With(prometheus.Labels{"server": server}).Inc()
}
