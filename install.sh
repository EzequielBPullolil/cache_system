#!/bin/bash

# Installation script for TCP_SERVER and cachex

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install it before running this script."
    exit 1
fi

# Compile and install TCP_SERVER
echo "Compiling and installing TCP_SERVER..."
go build -o ./tmp/cache_server ./tcp_server/main.go
if [ $? -eq 0 ]; then
    DIRECTORY="$HOME/.cache_system"
    mkdir -p $DIRECTORY
    mv ./tmp/cache_server $DIRECTORY
else
    echo "Error: Failed to compile TCP_SERVER. Check the previous errors."
    exit 1
fi

# Compile and install cachex
echo "Compiling and installing cachex..."
go build -o ./tmp/cachex ./shell/main.go
if [ $? -eq 0 ]; then
    mv ./tmp/cachex $DIRECTORY
else
    echo "Error: Failed to compile cachex. Check the previous errors."
    exit 1
fi

# Update PATH in ~/.bashrc and ~/.profile
echo "Updating PATH in ~/.bashrc and ~/.profile..."
echo "export PATH=\$PATH:$DIRECTORY" >> ~/.bashrc
echo "export PATH=\$PATH:$DIRECTORY" >> ~/.profile

echo "Installation completed successfully."
