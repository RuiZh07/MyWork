# NFC_Tag_UPoint
 This project aim to provide MS&T a convenient way of sharing personal contact information, such as Linkdin, Snapchat, Instagram, etc.
<br>

Requirement: Golang 1.19, postgreSQL

**Run following command to install all required software for first time** <br>
*this will start the web server as well*

`sudo bash ./setup.sh`
<br><br>

Run following command to start the server in the future <br>
`go run main.go`

--------

TODO

- ~Setup PostgreDB to the website~
- ~Finish sign up function (encrypt pw, choosing university)~
- ~Test log in~
- ~Create university email domain database~
- ~Let user select it's university and match it's email domain with database. Return error if mismatched~
- ~Create user dashboard~
- Redesign database schemas 
- Apply redeisgned shemas to database
- ~Add user verifyChecker to check if redirect the user~
- ~Ability to create new profile pages with links~
- ~Update the json files of social media~
- ~Personal page within links~
- ~Assign TagID to NFC tag~
- ~Function to link TagID to user's page~
- UI/UX update (Ruiz)
