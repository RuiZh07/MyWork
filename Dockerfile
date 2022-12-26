FROM golang:latest

# install PostgreSQL and create a user and database for the app
RUN apt-get update && apt-get install -y postgresql postgresql-contrib
RUN service postgresql start && \
    su postgres -c "psql -c \"CREATE USER cyw WITH PASSWORD 'cyw';\"" && \
    su postgres -c "psql -c \"CREATE DATABASE wacave WITH OWNER cyw;\""

# set the working directory for the app
WORKDIR /app

# copy the source code for the app
COPY . .

# build the app
RUN go build -o main .

# create the tables in the database
RUN service postgresql start && \
    su postgres -c "psql -d wacave -c \"CREATE TABLE users (email TEXT PRIMARY KEY, password TEXT, university TEXT);\"" && \
    su postgres -c "psql -d wacave -c \"CREATE TABLE universities (name TEXT PRIMARY KEY, domain TEXT, city TEXT, state TEXT);\""

# run the fetchData script to populate the universities table
RUN service postgresql start && go run src/fetchData.go

# expose the port for the app
EXPOSE 8080

# run the app and the PostgreSQL server
CMD service postgresql start && ./main

