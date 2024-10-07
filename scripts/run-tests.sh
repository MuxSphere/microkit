#!/bin/bash

# Run tests for all services otherwise go to individual services and run go test -v
for service in api-gateway service-a service-b; do
    echo "Running tests for $service"
    cd $service
    go test -v ./...
    cd ..
done

echo "All tests completed"