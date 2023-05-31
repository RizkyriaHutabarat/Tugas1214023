package namapackage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"os"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProfile(phone string, almt string, log string, biodata People) (InsertedID interface{}) {
	var profile Profile
	profile.Phone_number = phone
	profile.Alamat = almt
	profile.Login = log
	profile.Biodata = biodata
	return InsertOneDoc("tugas_db", "profile", profile)
}

func GetPeopleFromPhoneNumber(phone_number string) (people Profile) {
	People := MongoConnect("tugas_db").Collection("profile")
	filter := bson.M{"phone_number": phone_number}
	err := People.FindOne(context.TODO(), filter).Decode(&people)
	if err != nil {
		fmt.Printf("GetPeopleFromPhoneNumber: %v\n", err)
	}
	return people
}
