package Initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func ConnectToMongoDB() {
	// Load MongoDB URI from environment variable
	mongoURI := os.Getenv("DB_MONGO_URL")
	dbName := os.Getenv("MONGO_DB_NAME")

	if mongoURI == "" || dbName == "" {
		log.Fatal("Missing MONGO_URI or MONGO_DB_NAME in environment")
	}

	// Create a new client and set connection options
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Ping to ensure the connection is established
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Assign to global variables
	MongoClient = client
	MongoDatabase = client.Database(dbName)
}

func LoadEnvironmentVariables() {
	//Load environment variable
	var err error = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
