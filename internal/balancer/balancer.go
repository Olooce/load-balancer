package balancer

import (
	"load-balancer/internal/config"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	config  *config.LoadBalancerConfig
	mu      sync.Mutex
	servers []*Server
}

func NewLoadBalancer(cfg *config.LoadBalancerConfig) *LoadBalancer {
	lb := &LoadBalancer{
		config:  cfg,
		servers: []*Server{},
	}
	for _, srvCfg := range cfg.Servers {
		lb.servers = append(lb.servers, NewServer(srvCfg.Address, srvCfg.Weight))
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
	// load balancing logic
}
