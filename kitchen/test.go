package kitchen

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)



func ConnectToDB () *mongo.Client {
	// Create a new client and connect to the server
	err := godotenv.Load(".env")
	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	uri := os.Getenv("mongoURI") // default uri is local
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	 //client.Disconnect(context.TODO())
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Successfully connected and pinged.")
	return client
}