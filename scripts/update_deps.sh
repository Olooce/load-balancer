#!/bin/bash

# Navigate to the project root directory
cd "$(dirname "$0")/.."

# Update all dependencies to their latest versions
echo "Updating dependencies to the latest versions..."
go get -u ./...

# Clean up any unused dependencies
echo "Tidying up go.mod and go.sum files..."
go mod tidy

# Verify the updated dependencies
echo "Verifying dependencies..."
go mod verify

# Optionally, run tests to ensure everything works correctly
echo "Running tests..."
go test ./...

echo "Dependencies have been updated and verified."
