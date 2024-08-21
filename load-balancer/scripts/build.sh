#!/bin/bash
echo "Building Load Balancer..."
cd ../cmd/load-balancer
go build -o load-balancer
echo "Build complete."
