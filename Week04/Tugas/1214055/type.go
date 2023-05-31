package tugas

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func InsertPresensi(lokasi string, phonenumber string, checkin string, biodata Peneliti) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Location = lokasi
	presensi.Phone_number = phonenumber
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc("tugas4_db", "presensi", presensi)
}

func GetPenelitiFromPhoneNumber(phone_number string) (staf Jadwal) {
	peneliti := MongoConnect("tugas4_db").Collection("presensi")
	filter := bson.M{"phone_number": phone_number}
	err := peneliti.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getPenelitiFromPhoneNumber: %v\n", err)
	}
	return staf
}
