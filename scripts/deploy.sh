#!/bin/bash

# Get the name of the current directory
current_directory=$(basename "$(pwd)")

# Check if the current directory is not "docker-to-kubernetes-example"
if [ "$current_directory" != "docker-to-kubernetes-example" ]; then
    echo "Please run this from the root directory of the project (docker-to-kubernetes-example)."
    exit 1
fi

echo "Deploying the application..."
kubectl apply -f server/deployment.yaml
kubectl apply -f server/service.yaml
kubectl apply -f storage/redis-deployment.yaml
kubectl apply -f storage/redis-service.yaml