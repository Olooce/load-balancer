package monitor

import (
    "fmt"
    "html/template"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus metrics
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

// Template for the live-updating HTML page
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Load Balancer Monitoring</title>
    <style>
        body { font-family: Arial, sans-serif; }
        table { width: 100%; border-collapse: collapse; margin: 20px 0; }
        th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
        th { background-color: #f4f4f4; }
        h1 { margin: 20px 0; }
        .refresh { margin: 20px 0; }
    </style>
    <script>
        function refreshPage() {
            window.location.reload();
        }
        setInterval(refreshPage, 300); 
    </script>
</head>
<body>
    <h1>Load Balancer Monitoring</h1>
    <div class="refresh">
        <button onclick="refreshPage()">Refresh Now</button>
    </div>
    <h2>Server Metrics</h2>
    <table>
        <thead>
            <tr>
                <th>Server</th>
                <th>Request Count</th>
            </tr>
        </thead>
        <tbody>
            {{range .Servers}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.RequestCount}}</td>
            </tr>
            {{end}}
        </tbody>
    </table>
</body>
</html>
`

type ServerData struct {
	Name         string
	RequestCount int
}

type PageData struct {
	Servers []ServerData
}

func Start(port int) {
	// Register Prometheus metrics
	prometheus.MustRegister(requestCount)

	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		// Create a new Prometheus registry
		registry := prometheus.NewRegistry()
		registry.MustRegister(requestCount)

		// Gather the metrics
		metrics, err := registry.Gather()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Extract and parse the metrics
		var data PageData
		for _, mf := range metrics {
			for _, m := range mf.GetMetric() {
				labels := m.GetLabel()
				if len(labels) > 0 {
					serverName := labels[0].GetValue()
					value := m.GetCounter().GetValue()
					fmt.Printf("Parsed server: %s with request count: %d\n", serverName, int(value)) // Debug output
					data.Servers = append(data.Servers, ServerData{
						Name:         serverName,
						RequestCount: int(value),
					})
				}
			}
		}

		tmpl, err := template.New("monitor").Parse(htmlTemplate)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Expose Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Starting server on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {
	Start(8080)

	// Example of how to increment metrics
	go func() {
		for {
			requestCount.With(prometheus.Labels{"server": "server1"}).Inc()
			requestCount.With(prometheus.Labels{"server": "server2"}).Inc()
			requestCount.With(prometheus.Labels{"server": "server3"}).Inc()
		}
	}()
}
