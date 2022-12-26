package main

import (
	"log"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
	"NFC_Tag_UPoint/controller"
)

func main(){

	// log.Println("Waitting for server to boot in 10s")
	// time.Sleep(10 * time.Second)

	// Create table in database
	// createTable()

	// Initialize standard go html template engine
	engine := html.New("./templates", ".html")
	
	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/", index)
	app.Get("/signup", controller.LoadRegister)
	app.Get("/login", controller.LoadLoginPage)
	app.Post("/register", controller.HandleRegistration)
	app.Post("/login", controller.HandleLogin)



	// Start server
	log.Fatal(app.Listen(":8080"))
}

func index(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", nil)
}

func createTable(){
	// Connect to the PostgreSQL server.
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/wacave?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the universities table.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS universities (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			domain VARCHAR(255) NOT NULL,
			city VARCHAR(255) NOT NULL,
			state VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

