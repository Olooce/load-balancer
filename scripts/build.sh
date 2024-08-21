#!/bin/bash

echo "Building Load Balancer..."
cd ./cmd/load-balancer || { echo "Directory not found"; exit 1; }
go build -o load-balancer
echo "Build complete."
