#!/bin/bash

# Install golang
echo "Installing required package"
sudo apt-get update
sudo apt install snapd
snap install go --classic

# Install postgreSQL
echo "Installing postgreSQL"
sudo apt-get update
sudo apt install postgresql postgresql-contrib postgresql-client

# Starting database
echo "Starting Database"
sudo service postgresql start

# Creating user "admin" with password "admin"
echo "Creating user admin with password 'admin' "
sudo -u postgres psql -c "CREATE USER admin WITH PASSWORD 'admin';"

# Creating database "wacave"
echo "Creating database 'wacave'"
sudo -u postgres psql -c "CREATE DATABASE wacave;"


# Grant all permission to "admin" for "wacave" database
echo "Granting all permissions to user 'admin' for database 'wacave'"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;"

# Starting web service
echo "Starting web server"
go mod tidy
go run main.go