package tests

import (
	"load-balancer/internal/balancer"
	"load-balancer/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLoadBalancer tests the load balancer's routing functionality
func TestLoadBalancer(t *testing.T) {
	// Define mock servers with responses
	mockServer1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server 1 Response"))
	}))
	defer mockServer1.Close()

	mockServer2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server 2 Response"))
	}))
	defer mockServer2.Close()

	// Create LoadBalancer configuration
	cfg := &config.LoadBalancerConfig{
		Servers: []config.ServerConfig{
			{Address: mockServer1.URL, Weight: 1},
			{Address: mockServer2.URL, Weight: 2},
		},
		LoadBalancingAlgorithm: "round_robin",
	}

	// Initialize LoadBalancer
	lb := balancer.NewLoadBalancer(cfg)

	// Start the LoadBalancer in a goroutine
	go lb.Start()

	// Allow some time for the server to start
	// Replace this with a more robust way of checking server readiness if needed
	//time.Sleep(time.Second)

	// Create a test HTTP client to send requests to the LoadBalancer
	client := &http.Client{}

	// Test case 1: Check if the load balancer forwards requests correctly
	t.Run("TestLoadBalancing", func(t *testing.T) {
		resp, err := client.Get("http://localhost:8080") // Adjust port based on how `Start` is configured
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		expected := "Server 1 Response" // Expected response based on round-robin
		if got := readBody(resp); got != expected {
			t.Errorf("Expected response %s but got %s", expected, got)
		}
	})

	// Test case 2: Check if requests are distributed in a round-robin manner
	t.Run("TestRoundRobin", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			resp, err := client.Get("http://localhost:8080") // Adjust port based on how `Start` is configured
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			expected := "Server 1 Response"
			if i%2 != 0 {
				expected = "Server 2 Response"
			}
			if got := readBody(resp); got != expected {
				t.Errorf("Expected response %s but got %s", expected, got)
			}
		}
	})
}

// Helper function to read the body of the HTTP response
func readBody(resp *http.Response) string {
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	return string(body[:n])
}
