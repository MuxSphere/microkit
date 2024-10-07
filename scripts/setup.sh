#!/bin/bash

set -e

echo "Setting up MicroKit your reliable Microservices Starter Template..."

# Check if Docker is installed
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Please install Docker and try again."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null
then
    echo "Docker Compose is not installed. Please install Docker Compose and try again."
    exit 1
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# Build Docker images
echo "Building Docker images..."
docker-compose build

# Set up local development environment
echo "Setting up local development environment..."
cp .env.example .env
echo "Please edit the .env file with your specific configuration."

# Create necessary directories
mkdir -p data/postgres

echo "Setup complete! You can now start the services with 'docker-compose up'"
