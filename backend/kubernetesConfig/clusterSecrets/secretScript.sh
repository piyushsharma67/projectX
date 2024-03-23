#!/bin/bash

# Get the directory of the script
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")

# Input file location
input_file="$SCRIPT_DIR/globalConstants.txt"

# Check if input file exists
if [ ! -f "$input_file" ]; then
    echo "Error: Input file '$input_file' not found"
    exit 1
fi

# Read the content of the input file
content=$(cat "$input_file")

# Encode the content using base64
encoded_content=$(echo -n "$content" | base64)

# Create the Secret YAML file
cat <<EOF > secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: cluster-global-constants
type: Opaque
data:
  secret-data: $encoded_content
EOF

echo "Secret YAML file 'secret.yaml' created successfully"

echo "Deploying the Secret..."
kubectl apply -f secret.yaml

echo "Secret deployed successfully"