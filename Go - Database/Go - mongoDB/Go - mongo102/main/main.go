package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Database0:Data123@cluster0.ufgosci.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx) // Discoonnect from the database

	// Create DataBase
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	// CreateCollection(ctx, podcastsCollection)
	FindBreed(ctx, podcastsCollection, "Latte")
	// DeleteCollection(ctx, podcastsCollection, "Latte")

}

func DeleteCollection(ctx context.Context, client *mongo.Collection, dbName string) error {

	Filter := bson.M{"name": dbName}
	Fcursor, err := client.Find(ctx, Filter)
	if err != nil {
		log.Fatal(err)
	}
	// Find specific DataBase
	var vans []bson.M
	if err = Fcursor.All(ctx, &vans); err != nil {
		log.Fatal(err)
	}

	return err
}

func CreateCollection(ctx context.Context, client *mongo.Collection) {
	users := bson.D{
		{Name: "name", Value: "Momo"},
		{Name: "breed", Value: "Fish cat"},
	}

	bsonBytes, err := bson.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	podcastResult, err := client.InsertOne(ctx, bsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", podcastResult.InsertedID)
}

func FindBreed(ctx context.Context, client *mongo.Collection, dbName string) {
	// Find One DataBase
	var cat bson.M
	if err := client.FindOne(ctx, bson.M{}).Decode(&cat); err != nil {
		log.Fatal(err)
	}

	Filter := bson.M{"name": dbName}
	Fcursor, err := client.Find(ctx, Filter)
	if err != nil {
		log.Fatal(err)
	}
	// Find specific DataBase
	var vans []bson.M
	if err = Fcursor.All(ctx, &vans); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Turkish Vans : ", vans)
}
