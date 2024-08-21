package main

import (
	"load-balancer/internal/balancer"
	"load-balancer/internal/config"
	"load-balancer/internal/monitor"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("../../configs/load-balancer.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the load balancer
	lb := balancer.NewLoadBalancer(cfg)

	// Start monitoring if enabled
	if cfg.Monitor.Enabled {
		go monitor.Start(cfg.Monitor.Port)
	}

	// Start the load balancer
	lb.Start()
}
