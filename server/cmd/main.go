package main

import (
	"context"
	"fmt"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/coolcar"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("coolcar").Collection("account")
	// insertRows(ctx, collection)
	findRows(ctx, collection)
}

func findRows(c context.Context, col *mongo.Collection) {
	res := col.FindOne(c, bson.M{
		"open_id": "123",
	})
	fmt.Printf("%+v\n", res)
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		OpenId string             `bson:"open_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", row)
}

func insertRows(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(context.Background(), []interface{}{
		bson.M{"open_id": "qwe"},
		bson.M{"open_id": "asd"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
