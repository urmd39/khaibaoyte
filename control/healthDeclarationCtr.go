package control

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHealthDeclaration(idStr string) (hd bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("health_declaration")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hd); err != nil {
		log.Fatal(err)
	}
	return hd
}

func AddHealthDeclaration(hd bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("health_declaration")

	addResult, err := collection.InsertOne(ctx, hd)
	if err != nil {
		log.Fatal(err)
	}
	hd["_id"] = addResult.InsertedID
	return hd
}

func UpdateHealthDeclaration(idStr string, hd bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("health_declaration")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	collection.FindOneAndReplace(ctx, bson.M{"_id": id}, hd)
	hd["_id"] = id
	return hd
}

func AddQuestion(q bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("question_answer")

	addResult, err := collection.InsertOne(ctx, q)
	if err != nil {
		log.Fatal(err)
	}
	q["_id"] = addResult.InsertedID
	return q
}

func GetQuestion(idStr string) (q bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("question_answer")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&q); err != nil {
		log.Fatal(err)
	}
	return q
}

func UpdateQuestion(idStr string, q bson.M) bson.M {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("question_answer")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	collection.FindOneAndReplace(ctx, bson.M{"_id": id}, q)
	q["_id"] = id
	return q
}
