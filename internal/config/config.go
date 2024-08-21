package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ServerConfig represents the configuration for each backend server
type ServerConfig struct {
	Address string `yaml:"address"`
	Weight  int    `yaml:"weight"`
}

// ScalingConfig represents the scaling configuration for the load balancer
type ScalingConfig struct {
	Mode       string `yaml:"mode"`
	MaxServers int    `yaml:"max_servers"`
	MinServers int    `yaml:"min_servers"`
}

// MonitorConfig represents the monitoring configuration
type MonitorConfig struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

// LoadBalancerConfig represents the overall configuration for the load balancer
type LoadBalancerConfig struct {
	Servers                []ServerConfig `yaml:"servers"`
	LoadBalancingAlgorithm string         `yaml:"load_balancing_algorithm"`
	Scaling                ScalingConfig  `yaml:"scaling"`
	Monitor                MonitorConfig  `yaml:"monitor"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(path string) (*LoadBalancerConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return nil, err
	}
	defer file.Close()

	var config LoadBalancerConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	log.Printf("Loaded config: %+v", config)
	return &config, nil
}
