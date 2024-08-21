package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Address string `yaml:"address"`
	Weight  int    `yaml:"weight"`
}

type LoadBalancerConfig struct {
	Servers                []ServerConfig `yaml:"servers"`
	LoadBalancingAlgorithm string         `yaml:"load_balancing_algorithm"`
	Scaling                struct {
		Mode       string `yaml:"mode"`
		MaxServers int    `yaml:"max_servers"`
		MinServers int    `yaml:"min_servers"`
	} `yaml:"scaling"`
	Monitor struct {
		Enabled bool `yaml:"enabled"`
		Port    int  `yaml:"port"`
	} `yaml:"monitor"`
}

func LoadConfig(file string) (*LoadBalancerConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config LoadBalancerConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func init() {
	config, err := LoadConfig("configs/load-balancer.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Loaded config: %+v\n", config)
}
