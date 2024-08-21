#!/bin/bash

# Get the root directory of the project
PROJECT_ROOT=$(dirname "$(realpath "$0")")/..

# Start the example servers
echo "Starting Server 1..."
go run "$PROJECT_ROOT/examples/server1/main.go" &
SERVER1_PID=$!

echo "Starting Server 2..."
go run "$PROJECT_ROOT/examples/server2/main.go" &
SERVER2_PID=$!

# Start the load balancer
echo "Starting Load Balancer..."
go run "$PROJECT_ROOT/cmd/load-balancer/main.go" &
LOAD_BALANCER_PID=$!

# Wait for all processes to finish
wait $SERVER1_PID $SERVER2_PID $LOAD_BALANCER_PID
