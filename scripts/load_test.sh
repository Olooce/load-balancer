#!/bin/bash

# Default parameters
NUM_CALLS=1000
VERBOSITY=0
URL="http://localhost:8080"  # Default URL of the load balancer
WAIT_INTERVAL=1000  # Default wait interval after every 1000 requests

# Parse command line arguments
while getopts n:v:u:w: flag
do
    case "${flag}" in
        n) NUM_CALLS=${OPTARG};;
        v) VERBOSITY=${OPTARG};;
        u) URL=${OPTARG};;
        w) WAIT_INTERVAL=${OPTARG};;
    esac
done

# Function to make a single call to the load balancer
make_request() {
    if [ $VERBOSITY -gt 0 ]; then
        curl -s -o /dev/null -w "%{http_code} %{url_effective}\n" "$URL"
    else
        curl -s -o /dev/null "$URL"
    fi
}

# Execute the load test
for (( i=1; i<=NUM_CALLS; i++ ))
do
    make_request &
    if (( i % 1000 == 0 )); then
        if [ $WAIT_INTERVAL -gt 0 ]; then
            sleep $WAIT_INTERVAL  # Wait if WAIT_INTERVAL is greater than 0
        fi
        wait  # Wait for all background jobs to finish every 1000 requests
    fi
done

# Wait for any remaining background jobs
wait

echo "Completed $NUM_CALLS requests to $URL"
