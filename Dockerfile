FROM golang:latest

# install PostgreSQL and create a user and database for the app
RUN apt-get update && apt-get install -y postgresql postgresql-contrib
RUN service postgresql start && \
    su postgres -c "psql -c \"CREATE USER admin WITH PASSWORD 'admin';\"" && \
    su postgres -c "psql -c \"CREATE DATABASE wacave WITH OWNER admin;\"" && \
    su postgres -c "psql -d wacave -c \"CREATE TABLE users (id serial PRIMARY KEY, email text NOT NULL, password text NOT NULL, university text NOT NULL);\"" && \
    su postgres -c "psql -d wacave -c \"CREATE TABLE universities (name VARCHAR(255) NOT NULL, domain VARCHAR(255) NOT NULL, city VARCHAR(255) NOT NULL, state VARCHAR(255) NOT NULL);\"" && \
    su postgres -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;\""

# set the working directory for the app
WORKDIR /app

# copy the source code and the university data file for the app
COPY . .
COPY data/universityData.json .

# build the app
RUN go build -o main .

# create the tables in the database
# RUN service postgresql start && \
#     su postgres -c "psql -d wacave -c \"CREATE TABLE users (id serial PRIMARY KEY, email text NOT NULL, password text NOT NULL, university text NOT NULL);\"" && \
#     su postgres -c "psql -d wacave -c \"CREATE TABLE universities (name VARCHAR(255) NOT NULL, domain VARCHAR(255) NOT NULL, city VARCHAR(255) NOT NULL, state VARCHAR(255) NOT NULL);\""


# expose the port for the app
EXPOSE 8080

# && go run src/fetchData.go

# run the app and the PostgreSQL server and run fetchData.go to store all universities data in the database
CMD service postgresql start && go run data/fetchData.go && ./main
