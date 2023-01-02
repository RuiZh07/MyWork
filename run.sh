#!/bin/bash

# Build the Go web server image
docker build -t go-server .

# Start the Docker containers using docker-compose
docker compose up -d

# Stop containers
# docker compose down