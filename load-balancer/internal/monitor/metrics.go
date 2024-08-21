package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "load_balancer_requests_total",
			Help: "Total number of requests received by the load balancer",
		},
		[]string{"server"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

func RecordRequest(server string) {
	requestCount.With(prometheus.Labels{"server": server}).Inc()
}
