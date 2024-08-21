package balancer

import (
	"load-balancer/internal/config"
	"load-balancer/internal/server"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	config  *config.LoadBalancerConfig
	mu      sync.Mutex
	servers []*server.Server
	next    int // Index for round-robin
}

func NewLoadBalancer(cfg *config.LoadBalancerConfig) *LoadBalancer {
	lb := &LoadBalancer{
		config:  cfg,
		servers: []*server.Server{},
		next:    0, // Initialize round-robin index
	}
	for _, srvCfg := range cfg.Servers {
		lb.servers = append(lb.servers, server.NewServer(srvCfg.Address, srvCfg.Weight))
	}
	return lb
}

func (lb *LoadBalancer) Start() {
	http.HandleFunc("/", lb.handleRequest)
	http.ListenAndServe(":8080", nil)
}

func (lb *LoadBalancer) handleRequest(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Select server using round-robin method
	server := lb.servers[lb.next]
	lb.next = (lb.next + 1) % len(lb.servers) // Update index for next request

	server.HandleRequest(w, r)
}
