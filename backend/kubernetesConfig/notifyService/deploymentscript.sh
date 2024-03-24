#!/bin/bash

# Define the paths to your service and deployment YAML files
SERVICE_YAML="./service.yaml"
DEPLOYMENT_YAML="./deployment.yaml"

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
    echo "kubectl is not installed. Please install it before running this script."
    exit 1
fi

# Check if service YAML file exists
if [ ! -f "$SERVICE_YAML" ]; then
    echo "Service YAML file not found at: $SERVICE_YAML"
    exit 1
fi

# Check if deployment YAML file exists
if [ ! -f "$DEPLOYMENT_YAML" ]; then
    echo "Deployment YAML file not found at: $DEPLOYMENT_YAML"
    exit 1
fi

# Apply the service YAML
echo "Applying service YAML..."
kubectl apply -f "$SERVICE_YAML"

# Apply the deployment YAML
echo "Applying deployment YAML..."
kubectl apply -f "$DEPLOYMENT_YAML"

# Check if deployment was successful
if [ $? -eq 0 ]; then
    echo "Deployment successful."
else
    echo "Error deploying service and deployment."
fi
