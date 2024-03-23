#!/bin/bash

# Define the path to the config directory
CONFIG_DIR="."

# Loop through each subdirectory of the config directory
for dir in "$CONFIG_DIR"/*; do
    # Check if the item is a directory
    if [ -d "$dir" ]; then
        # Navigate to the directory
        cd "$dir" || { echo "Error: Unable to change directory to $dir"; exit 1; }
        
        # Get a list of script files in the directory
        scripts=$(find . -maxdepth 1 -type f -name "*.sh")

        # Loop through each script file
        for script in $scripts; do
            # Execute the script
            echo "Executing script: $script"
            bash "$script"
            echo "Script execution finished: $script"
        done

        # Navigate back to the original directory
        cd - > /dev/null || { echo "Error: Unable to change directory back to the original directory"; exit 1; }
    fi
done
