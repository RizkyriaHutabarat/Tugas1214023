package namapackage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

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

func InsertKuesioner(long float64,lat float64, lokasi string, email string, status string, biodata Responden) (InsertedID interface{}) {
	var kuesioner Kuesioner
	kuesioner.Latitude = long
	kuesioner.Longitude = lat
	kuesioner.Location = lokasi
	kuesioner.Email = email
	kuesioner.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	kuesioner.Status = status
	kuesioner.Biodata = biodata
	return InsertOneDoc("data_db", "kuesioner", kuesioner)
}

func GetRespondenFromEmail(email string) (responden Kuesioner) {
	Responden := MongoConnect("data_db").Collection("kuesioner")
	filter := bson.M{"email": email}
	err := Responden.FindOne(context.TODO(), filter).Decode(&responden)
	if err != nil {
		fmt.Printf("getRespondenFromEmail: %v\n", err)
	}
	return responden
}
