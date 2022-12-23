//mongo db
package models

import(
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

// MongoInstance contains the Mongo client and database object
type MongoInstance struct {
	Client *mongo.Client
	Db *mongo.Database
}

var mg MongoInstance

// Database settings (inserted own database name and connection URI)
const dbName = "wacave"
const mongoURI = "mongodb://user:password@localhost:27017/" + dbName

// User Struct
type User struct {
	Name string
	Email string
	Password string
	University string
}

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db: db,
	}

	return nil
}