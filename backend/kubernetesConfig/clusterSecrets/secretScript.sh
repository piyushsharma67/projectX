#!/bin/bash

# Check if input file is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <input_file>"
    exit 1
fi

# Input file
input_file="$1"

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
  name: clusterGlobalConstants
type: Opaque
data:
  secret-data: $encoded_content
EOF

echo "Secret YAML file 'secret.yaml' created successfully"