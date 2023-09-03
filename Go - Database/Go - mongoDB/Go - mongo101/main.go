package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Insert Document in the DataBase
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Insert One DataBase
	demoDB := client.Database("demo")
	err = demoDB.CreateCollection(ctx, "cats")
	if err != nil {
		log.Fatal(err)
	}
	catsCollection := demoDB.Collection("cats")
	defer catsCollection.Drop(ctx)
	result, err := catsCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Mocha"},
		{Key: "breed", Value: "Turkish Van"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", result)

	// Insert many Database
	manyResult, err := catsCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "name", Value: "Latte"},
			{Key: "breed", Value: "Maine Coon"},
		},
		bson.D{
			{Key: "name", Value: "Trouble"},
			{Key: "breed", Value: "Domestic Shorthair"},
		}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", manyResult)

	// Find the DataBase
	/*
		cursor, err := catsCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var cats []bson.M
		if err = cursor.All(ctx, &cats); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Cats :", cats)
	*/

	// Label all Find the DataBase
	cursor, err := catsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var kitty bson.M
		if err = cursor.Decode(&kitty); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Kitty : ", kitty)
	}

	// Find One DataBase
	var cat bson.M
	if err = catsCollection.FindOne(ctx, bson.M{}).Decode(&cat); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cat : ", cat)
	Filter := bson.M{"breed": "Turkish Van"}
	Fcursor, err := catsCollection.Find(ctx, Filter)
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
