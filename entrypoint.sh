#!/bin/bash

# Wait for the PostgreSQL server to start
echo "Waiting for PostgreSQL to start..."
while ! pg_isready -q -h $PGHOST -p $PGPORT -U $PGUSER
do
  echo "."
  sleep 1
done

# Create the 'wacave' database
echo "Creating 'wacave' database..."
psql -v ON_ERROR_STOP=1 --username "$PGUSER" <<-EOSQL
    CREATE DATABASE wacave;
EOSQL

# Create the 'admin' user
echo "Creating 'admin' user..."
psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "wacave" <<-EOSQL
    CREATE USER admin WITH PASSWORD 'admin';
EOSQL

# Grant privileges to the 'admin' user
echo "Granting privileges to 'admin' user..."
psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "wacave" <<-EOSQL
    GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;
EOSQL

# Create the 'users' table
echo "Creating 'users' table..."
psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "wacave" <<-EOSQL
    CREATE TABLE users (
        id serial PRIMARY KEY,
        email text NOT NULL,
        password text NOT NULL,
        university text NOT NULL
    );
EOSQL

# Create the 'universities' table
echo "Creating 'universities' table..."
psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "wacave" <<-EOSQL
    CREATE TABLE universities (
        name VARCHAR(255) NOT NULL,
        domain VARCHAR(255) NOT NULL,
        city VARCHAR(255) NOT NULL,
        state VARCHAR(255) NOT NULL
    );
EOSQL

# Run the fetchData.go script
echo "Importing data from 'universityData.json'..."
go run src/fetchData.go