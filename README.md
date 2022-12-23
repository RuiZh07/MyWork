# NFC_Tag_UPoint
 This project aim to provide MS&T a convinient way of sharing personal contact information, such as Linkdin, Snapchat, Instagram, etc.
<br>

 To start the server, run command <br>
`go mod tidy`

 then <br>
`go run main.go`

The default port is set to `8080`, therefore, open your browser and paste following url into your browser
`localhost:8080`

You should see the current design for the page


--------

TODO

- Setup MongoDB to the website
- Finish sign up function (encrypt pw, choosing university, check email domain)
- Test log in
- Add user verifyChecker to check if redirect the user
- Personal page within links
- Assign TagID to NFC tag
- Function to link TagID to user's page
- UI/UX update (Ruiz)
