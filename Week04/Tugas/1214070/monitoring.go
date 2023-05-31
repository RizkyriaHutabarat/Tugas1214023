package monitoring

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

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

func InsertOrangTua(nama_ot string,phone_number string, anak Mahasiswa) (InsertedID interface{}) {
	var orangtua OrangTua
	orangtua.Nama_OT = nama_ot
	orangtua.Phone_number = phone_number
	orangtua.Anak = anak
	return InsertOneDoc("monitoring_db", "orangtua", orangtua)
}

func GetTemaFromPhoneNumber(phone_number string) (nomor OrangTua) {
	Tema := MongoConnect("monitoring_db").Collection("orangtua")
	filter := bson.M{"phone_number": phone_number}
	err := Tema.FindOne(context.TODO(), filter).Decode(&nomor)
	if err != nil {
		fmt.Printf("getTemaFromPhoneNumber: %v\n", err)
	}
	return nomor
}
