#!/bin/bash

echo "Build Web Server Container"
docker build -t golang-web-server -f Dockerfile.web .

echo "Build Database Container"
docker build -t postgresql-server -f Dockerfile.db .

echo "Run WaCave network"
sudo docker compose up -d

# Wait for the postgresql container to start
echo "Waitting 5 Seconds For Server To Start"
sleep 5

# Run the setup_database.sh script
echo "Setting Up Database"
sudo bash ./setup_database.sh

sudo docker compose down