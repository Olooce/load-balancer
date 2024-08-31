package monitor

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    // Define a new counter vector with a label for the server name
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests processed, labeled by server.",
        },
        []string{"server"}, // Labels by server
    )
)
