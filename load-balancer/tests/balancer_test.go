package tests

import (
	"load-balancer/internal/balancer"
	"load-balancer/internal/config"
	"testing"
)

func TestLoadBalancer(t *testing.T) {
	cfg := &config.LoadBalancerConfig{
		Servers: []config.ServerConfig{
			{Address: "http://localhost:8081", Weight: 1},
			{Address: "http://localhost:8082", Weight: 2},
		},
		LoadBalancingAlgorithm: "round_robin",
	}

	lb := balancer.NewLoadBalancer(cfg)

	// Add test cases here to validate load balancing
}
