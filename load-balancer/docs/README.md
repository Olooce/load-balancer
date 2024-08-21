---

# Go Load Balancer

## Overview

This project is a Go-based Load Balancer that evenly distributes incoming client requests across multiple servers. The load balancer supports both horizontal and vertical scaling and includes a built-in monitoring system accessible through a web interface. The project also provides example server implementations to test and demonstrate the load balancer's functionality.

## Features

- **Load Balancing Algorithms**: Supports different load balancing algorithms such as Round Robin.
- **Scaling**: Configurable scaling modes (horizontal and vertical) to manage server resources effectively.
- **Monitoring**: Integrated web-based monitoring to track load balancer performance and server status.
- **Example Servers**: Simple Go HTTP servers included to simulate and test load balancing.

## Project Structure

```
load-balancer/
├── cmd/
│   └── load-balancer/
│       └── main.go           # Entry point for the load balancer
├── internal/
│   ├── balancer/
│   │   ├── balancer.go       # Core load balancing logic
│   │   └── algorithms.go     # Load balancing algorithms
│   ├── server/
│   │   └── server.go         # Server management logic
│   ├── monitor/
│   │   ├── metrics.go        # Metrics collection
│   │   └── web.go            # Web interface for monitoring
│   └── config/
│       └── config.go         # Configuration management
├── examples/
│   ├── server1/
│   │   └── main.go           # Example server 1
│   ├── server2/
│   │   └── main.go           # Example server 2
│   └── docker-compose.yml    # Docker Compose configuration for example servers
├── configs/
│   └── load-balancer.yaml    # Load balancer configuration file
├── deploy/
│   ├── Dockerfile            # Dockerfile for deploying the load balancer
│   └── kubernetes.yaml       # Kubernetes deployment configuration
├── docs/
│   └── README.md             # Project documentation (this file)
├── scripts/
│   └── install_deps.sh       # Script to install dependencies
└── tests/
    └── balancer_test.go      # Unit tests for the load balancer
```

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Olooce/load-balancer.git
   cd load-balancer
   ```

2. **Install dependencies**:

   Run the provided script to install all dependencies required by the project:

   ```bash
   ./scripts/install_deps.sh
   ```

3. **Build the Load Balancer**:

   ```bash
   cd cmd/load-balancer
   go build -o load-balancer
   ```

## Configuration

The load balancer is configured using a YAML file located at `configs/load-balancer.yaml`. Here is a sample configuration:

```yaml
servers:
  - address: "http://localhost:8081"
    weight: 1
  - address: "http://localhost:8082"
    weight: 2

load_balancing_algorithm: "round_robin"

scaling:
  mode: "horizontal"  # options: horizontal, vertical
  max_servers: 10
  min_servers: 2

monitor:
  enabled: true
  port: 9090
```

- **Servers**: List of backend servers with their addresses and weights.
- **Load Balancing Algorithm**: The algorithm used to distribute requests (e.g., `round_robin`).
- **Scaling**: Configuration for scaling behavior, including mode, maximum, and minimum servers.
- **Monitor**: Monitoring settings, including whether monitoring is enabled and the port for the web interface.

## Running the Load Balancer

1. **Start the Load Balancer**:

   ```bash
   ./cmd/load-balancer/load-balancer
   ```

   The load balancer will start using the configuration specified in `configs/load-balancer.yaml`.

2. **Run Example Servers**:

   Open separate terminal windows or use a process manager to start the example servers:

   ```bash
   go run examples/server1/main.go
   go run examples/server2/main.go
   ```

   Alternatively, you can use Docker Compose:

   ```bash
   docker-compose -f examples/docker-compose.yml up
   ```

## Monitoring

If monitoring is enabled in the configuration, you can access the metrics and status of the load balancer by navigating to `http://localhost:9090/metrics` in your web browser.

## Testing

Unit tests for the load balancer are provided in the `tests` directory. Run the tests using:

```bash
go test ./tests/...
```

## Deployment

### Docker

To build and run the load balancer using Docker:

```bash
docker build -t load-balancer .
docker run -d -p 8080:8080 -p 9090:9090 load-balancer
```

### Kubernetes

Deploy the load balancer to a Kubernetes cluster using the provided `kubernetes.yaml` file:

```bash
kubectl apply -f deploy/kubernetes.yaml
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For any questions or support, please contact [oloostephen.dev@gmail.com].

---

