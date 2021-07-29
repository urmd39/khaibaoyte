package control

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
	ctx    context.Context
)

func AddPerson(p bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("person_info")

	addResult, err := collection.InsertOne(ctx, p)
	if err != nil {
		log.Fatal(err)
	}
	p["_id"] = addResult.InsertedID
	return p
}

func GetPerson(idStr string) (person bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("person_info")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&person); err != nil {
		log.Fatal(err)
	}
	return person
}

func UpdatePerson(idStr string, p bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("person_info")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	collection.FindOneAndReplace(ctx, bson.M{"_id": id}, p)
	p["_id"] = id
	return p
}
