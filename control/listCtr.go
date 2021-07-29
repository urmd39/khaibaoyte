package control

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetListSex() (list []bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("sex")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var sex bson.M
		if err = cursor.Decode(&sex); err != nil {
			log.Fatal(err)
		}
		list = append(list, sex)
	}
	return list
}

func GetListNationality() (list []bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("nationality")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var nationality bson.M
		if err = cursor.Decode(&nationality); err != nil {
			log.Fatal(err)
		}
		list = append(list, nationality)
	}
	return list
}

func GetListPerson() (list []bson.M) {
	client, ctx = Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("person_info")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var person bson.M
		if err = cursor.Decode(&person); err != nil {
			log.Fatal(err)
		}
		list = append(list, person)
	}
	return list
}

func GetListProvince() (list []*bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("province")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var province bson.M
		if err = cursor.Decode(&province); err != nil {
			log.Fatal(err)
		}
		list = append(list, &province)
	}
	return list
}

func GetListTown(province string) (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("town")

	cursor, err := collection.Find(ctx, bson.M{"province": province})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var town bson.M
		if err = cursor.Decode(&town); err != nil {
			log.Fatal(err)
		}
		list = append(list, town)
	}
	return list
}

func GetListVillage(town string) (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("village")

	cursor, err := collection.Find(ctx, bson.M{"town": town})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var village bson.M
		if err = cursor.Decode(&village); err != nil {
			log.Fatal(err)
		}
		list = append(list, village)
	}
	return list
}

func GetListQuestion() (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("question_answer")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var question bson.M
		if err = cursor.Decode(&question); err != nil {
			log.Fatal(err)
		}
		list = append(list, question)
	}
	return list
}

func GetListHealthDeclaration() (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("khaibaoyte").Collection("health_declaration")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var hd bson.M
		if err = cursor.Decode(&hd); err != nil {
			log.Fatal(err)
		}
		list = append(list, hd)
	}
	return list
}
