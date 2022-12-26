# NFC_Tag_UPoint
 This project aim to provide MS&T a convinient way of sharing personal contact information, such as Linkdin, Snapchat, Instagram, etc.
<br>

Requirement: Golang 1.19, postgreSQL

to install golang 1.19, you can run <br>
`sudo apt update` <br>
`sudo apt install snapd` <br>
`snap install go --classic` <br>

to install postgreSQL <br>
`sudo apt update` <br>
`sudo apt install postgresql postgresql-contrib` <br> <br>
start server with: 
`sudo service postgresql start` <br>
login to postgre server: 
`sudo -u postgres psql` <br>
create user with username "admin" and password "admin": 
`CREATE USER admin WITH PASSWORD 'admin';` <br>
create a database named "wacave": 
`CREATE DATABASE wacave;` <br>
give admin permission of the database:
`GRANT ALL PRIVILEGES ON DATABASE wacave TO admin;` <br>




 To start the server, run command <br>
`go mod tidy`
`go run main.go` 

The default port is set to `8080`, therefore, open your browser and paste following url into your browser
`localhost:8080`

You should see the current design for the page

--------

If you have docker and docker compose installed, you can run with command
`sudo bash run.sh`
to start the server without any effort!
(note: 12/25 the docker has permission issue when writing data to database, don't use it yet until we fix it)

--------

TODO

- ~Setup MongoDB to the website~
- ~Finish sign up function (encrypt pw, choosing university)~
- ~Test log in~
- Create university email domain database
- Let user select it's university and match it's email domain with database. Return error if mismatched
- Add user verifyChecker to check if redirect the user
- Personal page within links
- Assign TagID to NFC tag
- Function to link TagID to user's page
- UI/UX update (Ruiz)
