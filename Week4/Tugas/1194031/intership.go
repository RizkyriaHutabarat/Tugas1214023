package intership

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

func InsertIntership(npm string, kelas string, kampus string, judul string, biodata Mahasiswa) (InsertedID interface{}) {
	var intership Intership
	intership.Npm = npm
	intership.Kelas = kelas
	intership.Kampus = kampus
	intership.Judul = judul
	intership.Tanggal = primitive.NewDateTimeFromTime(time.Now().UTC())
	intership.Biodata = biodata
	return InsertOneDoc("tugas_db", "intership", intership)
}

func GetMahasiswaFromNama(nama string) (staf Intership) {
	mahasiswa := MongoConnect("tugas_db").Collection("intership")
	filter := bson.M{"nama": nama}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("Nama Mahasiswa : %v\n", err)
	}
	return staf
}