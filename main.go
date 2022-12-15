// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	//"fmt"
	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"
	"html/template"
	//"github.com/gin-gonic/gin"
)

var tmpl *template.Template
const signupPage = "templates/signup.html"

func main() {
	
	http.HandleFunc("/", foo)
	
	http.ListenAndServe(":8080", nil)
}

func init() {
    tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func foo(reswt http.ResponseWriter, req *http.Request) {
    tmpl.ExecuteTemplate(reswt, "index.html", nil)
}


