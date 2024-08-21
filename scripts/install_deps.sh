#!/bin/bash

# Start at the root of the project
PROJECT_ROOT=$(pwd)

# Find all Go module files in the project
GO_MODULE_FILES=$(find "$PROJECT_ROOT" -name "go.mod")

# Iterate over each Go module file and install dependencies
for MOD_FILE in $GO_MODULE_FILES; do
    MODULE_DIR=$(dirname "$MOD_FILE")
    echo "Processing module in $MODULE_DIR..."

    # Change to the module directory
    cd "$MODULE_DIR"

    # Install dependencies
    echo "Running go mod tidy to ensure all dependencies are correctly added..."
    go mod tidy

    echo "Running go get to install any missing dependencies..."
    go get -d ./...

    echo "Dependencies installed for $MODULE_DIR."
done

# Return to the project root directory
cd "$PROJECT_ROOT"

echo "All dependencies have been installed."
