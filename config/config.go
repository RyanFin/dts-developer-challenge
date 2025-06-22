package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func init() {
	// Load .env variables using Viper
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func ConnectDB() {
	uri := viper.GetString("MONGO_DB_ATLAS_URI")
	dbName := viper.GetString("MONGO_DB_NAME")

	fmt.Printf("%s -- %s", uri, dbName)

	if uri == "" || dbName == "" {
		log.Fatal("Missing MONGO_DB_ATLAS_URI or MONGO_DB_NAME in .env")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	TaskCollection = client.Database(dbName).Collection("tasks")
}
