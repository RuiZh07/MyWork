#!/bin/bash

# Start postgresql service
sudo service postgresql start

# Connect to postgresql as user "postgres"
sudo -u postgres psql

# Create a new user "admin" with password "admin"
CREATE USER admin WITH PASSWORD 'admin';

# Create a new database called "wacave"
CREATE DATABASE wacave;

# Grant all privileges on "wacave" to the "admin" user
GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;
