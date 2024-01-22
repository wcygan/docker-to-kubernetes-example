#!/bin/bash

# Get the pod name for your gRPC server
POD_NAME=$(kubectl get pods --selector=app=grpc-server -o jsonpath="{.items[0].metadata.name}")

# Set up port forwarding for your gRPC server
kubectl port-forward $POD_NAME 8080:8080