#!/bin/bash

# Connect to PostgreSQL as the postgres user
sudo -u postgres psql <<-EOSQL

# Create user with name 'admin' and password 'admin'
CREATE USER admin WITH PASSWORD 'admin';

# Create the database and user
CREATE DATABASE wacave;
GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;;

# Connect to the database as the user
\c wacave admin



# Create the users table
CREATE TABLE users (
    id serial PRIMARY KEY,
    email text NOT NULL,
    password text NOT NULL,
    university text NOT NULL
);

# Create the universities table
CREATE TABLE universities (
	name VARCHAR(255) NOT NULL,
	domain VARCHAR(255) NOT NULL,
	city VARCHAR(255) NOT NULL,
	state VARCHAR(255) NOT NULL
);

EOSQL
